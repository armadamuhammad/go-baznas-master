package lib

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/iancoleman/strcase"
)

// Response http response
type Response struct {
	Status           int          `json:"status"`                                           // http status
	Message          string       `json:"message"`                                          // response message
	ErrorDescription *string      `json:"error_description,omitempty" swaggerignore:"true"` // Error description
	ErrorData        *[]ErrorData `json:"error_data,omitempty" swaggerignore:"true"`        // Error fields
}

// ErrorData field error details
type ErrorData struct {
	// Namespace string      `json:"namespace" example:"User.Username"`                       // data namespace
	Name      string      `json:"name" example:"Username"`                                 // data name
	Path      string      `json:"path" example:"user.username"`                            // object property path
	Type      string      `json:"type,omitempty" example:"string"`                         // data type
	Value     interface{} `json:"value,omitempty" swaggertype:"string" example:"jane doe"` // value
	Validator string      `json:"validator" example:"required"`                            // validator type, see [more details](https://github.com/go-playground/validator#baked-in-validations)
	Criteria  interface{} `json:"criteria,omitempty" swaggertype:"number" example:"10"`    // criteria, example: if validator is gte (greater than) and criteria is 10, then it means a maximum of 10
	Message   string      `json:"-" example:"invalid value"`                               // Field message
}

type sqlPatterns struct {
	driver    string
	pattern   string
	validator string
}

var patterns = []sqlPatterns{
	{driver: "postgres", pattern: `ERROR:.*duplicate key.*"([^"]+)".*`, validator: "unique"},
	{driver: "mysql", pattern: `Duplicate entry.*for key '([^']+)`, validator: "unique"},
	{driver: "sqlite", pattern: `UNIQUE.*:\s*(.*)`, validator: "unique"},
	{driver: "postgres", pattern: `ERROR: null.*"([^"]+)".*`, validator: "required"},
	{driver: "sqlite", pattern: `.*NOT NULL.*:\s*(.*)`, validator: "required"},
}

var validateMessages = map[string]string{
	"required": "%v is required %v",
	"gte":      "%v must be greater than or equal %v",
	"gt":       "%v must be greater than %v",
	"lte":      "%v must be less than or equal %v",
	"lt":       "%v must be less than %v",
	"unique":   "%v already exists %v",
}

// Send response
func Send(c *fiber.Ctx, status int, responses ...interface{}) error {
	response := Response{
		Status: status,
	}
	for i := 0; i < len(responses); i++ {
		if e, ok := responses[i].(Response); ok {
			response = e
			response.Status = status
			break
		}
		if e, ok := responses[i].(error); ok {
			response.ErrorDescription = Strptr(e.Error())
			if errs, ok := e.(validator.ValidationErrors); ok {
				response.Status = 400
				response.Message = "Bad Request"
				response.ErrorDescription = Strptr("Request body does not meet the requirements")
				errorDetails := []ErrorData{}
				for _, err := range errs {
					fields := strings.Split(err.StructNamespace(), ".")
					field := strcase.ToSnake(err.Field())
					if len(fields) > 1 {
						names := []string{}
						for _, i := range fields[1:] {
							f := strcase.ToSnake(i)
							names = append(names, f)
						}
						field = strings.Join(names, ".")
					}
					fieldType := strings.ReplaceAll(fmt.Sprint(err.Type()), "*", "")
					regex := regexp.MustCompile(`.*\.([^\.]+)$`)
					fieldType = regex.ReplaceAllString(fieldType, "$1")
					errorData := ErrorData{
						Name:      strcase.ToSnake(err.Field()),
						Path:      field,
						Type:      fieldType,
						Value:     err.Value(),
						Validator: err.Tag(),
					}

					if i, e := strconv.Atoi(err.Param()); nil == e {
						errorData.Criteria = i
					} else {
						errorData.Criteria = err.Param()
					}

					if m, ok := validateMessages[errorData.Validator]; ok {
						errorData.Message = fmt.Sprintf(m, errorData.Name, errorData.Criteria)
					}
					errorDetails = append(errorDetails, errorData)
				}
				response.ErrorData = &errorDetails
				break
			}
		} else if e, ok := responses[i].(string); ok {
			response.Message = e
		}
	}

	if response.Message == "" && nil != response.ErrorDescription {
		response.Message = *response.ErrorDescription
	}

	return c.Status(status).JSON(response)
}

// ErrorBadRequest send http 400 bad request
//   message can contains string or error or nothing
//   example:
//   lib.ErrorBadRequest(c, "bad request")
//   lib.ErrorBadRequest(c, errors.New("bad request"))
//   lib.ErrorBadRequest(c) // default response message is Bad Request
func ErrorBadRequest(c *fiber.Ctx, message ...interface{}) error {
	if len(message) == 0 {
		message = append(message, "Bad request")
	}

	return Send(c, 400, message[0])
}

// ErrorNotFound send http 404 not found
func ErrorNotFound(c *fiber.Ctx, message ...string) error {
	if len(message) == 0 {
		message = append(message, "Not found")
	}

	return Send(c, 404, message[0])
}

// ErrorInternal send http 500 internal server error
func ErrorInternal(c *fiber.Ctx, message ...string) error {
	if len(message) == 0 {
		message = append(message, "Internal server error")
	}

	return Send(c, 500, message[0])
}

// ErrorConflict send http 409 conflict
//   message can contains string or error or nothing
//   example:
//   lib.ErrorConflict(c, "Conflict")
//   lib.ErrorConflict(c, errors.New("Conflict"))
//   lib.ErrorConflict(c) // default response message is Conflict
func ErrorConflict(c *fiber.Ctx, message ...interface{}) error {
	if len(message) == 0 {
		message = append(message, "Conflict")
	}

	responseMessage := "Conflict"
	if m, ok := message[0].(string); ok {
		responseMessage = m
	} else if m, ok := message[0].(error); ok {
		responseMessage = m.Error()
	}

	indexPrefixes := regexp.MustCompile(`^(unique|index|idx)_+`)

	for p := range patterns {
		pattern := patterns[p]
		re := regexp.MustCompile(pattern.pattern)
		if re.Match([]byte(responseMessage)) {
			fieldPath := re.ReplaceAllString(responseMessage, "$1")
			fieldPath = indexPrefixes.ReplaceAllString(fieldPath, "")
			fieldPath = strings.ReplaceAll(fieldPath, "__", ".")
			field := fieldPath
			fields := strings.Split(field, ".")
			if len(fields) > 1 {
				names := []string{}
				for _, i := range fields[1:] {
					names = append(names, strcase.ToSnake(i))
				}
				field = strings.Join(names, ".")
			}

			// detect field name based on gorm default index
			// set default gorm unique index example:
			// `gorm:"type:varchar(36);index:,unique,where:deleted_at is null;not null"`
			if !strings.Contains(fieldPath, ".") {
				splits := strings.Split(fieldPath, "_")
				splitLen := len(splits)
				if splitLen > 2 {
					splitIndex := (splitLen - 1) / 2
					prefix := strings.Join(splits[0:splitIndex], "_")
					suffixLen := len(prefix) + 1
					if len(fieldPath) > suffixLen {
						suffix := fieldPath[suffixLen:]
						if strings.HasPrefix(suffix, prefix) {
							field = suffix
							fieldPath = prefix + "." + suffix
						}
					}
				}
			}
			errorData := ErrorData{
				Name:      strcase.ToSnake(field),
				Path:      fieldPath,
				Validator: pattern.validator,
				Value:     nil,
			}

			if m, ok := validateMessages[pattern.validator]; ok {
				errorData.Message = fmt.Sprintf(m, errorData.Name, "")
			}

			response := Response{
				Message:          `Conflict`,
				ErrorDescription: Strptr(`Duplicate value`),
				ErrorData:        &[]ErrorData{errorData},
			}

			httpCode := 409

			// fallback to bad request if not unique validator
			if pattern.validator != "unique" {
				httpCode = 400
				response.Message = "Bad Request"
				response.ErrorDescription = Strptr(responseMessage)
			}

			return Send(c, httpCode, response)
		}
	}

	return Send(c, 409, responseMessage)
}

// OK send http 200 response
func OK(c *fiber.Ctx, result ...interface{}) error {
	if len(result) == 0 {
		result = append(result, Response{
			Status:  200,
			Message: "success",
		})
	}

	return c.Status(200).JSON(result[0])
}
