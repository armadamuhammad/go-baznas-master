//go:build !integration
// +build !integration

package controller

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

var (
	basePattern     = regexp.MustCompile(`(?i)@(accept|description|tags|summary|produce|param|router|success|failure)[ ]+(.*)`)
	filePattern     = regexp.MustCompile(`^(.*)_(get|put|post|delete|patch)\.go$`)
	commentPattern  = regexp.MustCompile(`^([^ ]+)[ ]+([^ ]+)[ ]+([^ ]+)[ ]+([^ ]+)[ ]+"?([^"]+)"?`)
	responsePattern = regexp.MustCompile(`^([^ ]+)[ ]+([^ ]+)[ ]+([^ ]+)([ "]+)?([^"]+)"?`)
	spacePattern    = regexp.MustCompile(`[ ]+`)
	methodPattern   = regexp.MustCompile(`[\[\]]+`)
	pathPattern     = regexp.MustCompile(`\{([^\}]+)\}`)
	pascalPattern   = regexp.MustCompile(`^(?i)(URL|ID|API)$`)
)

func getAllGoFiles(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	outputs := []string{}
	if nil == err {
		for f := range files {
			file := files[f]
			if file.IsDir() {
				outputs = append(outputs, getAllGoFiles(path.Join(dir, file.Name()))...)
			} else {
				if filePattern.MatchString(file.Name()) {
					outputs = append(outputs, path.Join(dir, file.Name()))
				}
			}
		}
	}

	return outputs
}

func toPascal(str string) string {
	pascalString := []string{}
	re := filePattern
	baseFileName := re.ReplaceAllString(filepath.Base(str), "${2}_${1}")
	baseFileNames := strings.Split(baseFileName, "_")
	for _, value := range baseFileNames {
		suffix := strings.ToLower(string(value[1:]))
		if pascalPattern.MatchString(value) {
			suffix = strings.ToUpper(string(value[1:]))
		}
		pascalString = append(pascalString, strings.ToUpper(string(value[0])))
		pascalString = append(pascalString, suffix)
	}

	return strings.Join(pascalString, "")
}

type SwaggerCommentParam struct {
	Name        string
	In          string
	Type        string
	Required    bool
	Description string
	Content     string
}

func (s *SwaggerCommentParam) ParseContent() {
	if s.Content != "" {
		re := commentPattern
		if re.MatchString(s.Content) {
			s.Name = re.ReplaceAllString(s.Content, "$1")
			s.In = re.ReplaceAllString(s.Content, "$2")
			s.Type = re.ReplaceAllString(s.Content, "$3")
			s.Required = re.ReplaceAllString(s.Content, "$4") == "true"
			s.Description = re.ReplaceAllString(s.Content, "$5")
		}
	}
}

type SwaggerCommentResponse struct {
	Status      string
	Code        int
	Type        string
	Model       string
	Description string
	Content     string
}

func (s *SwaggerCommentResponse) ParseContent() {
	if s.Content != "" {
		re := responsePattern
		if re.MatchString(s.Content) {
			codeInt, _ := strconv.Atoi(re.ReplaceAllString(s.Content, "$1"))
			s.Code = codeInt
			s.Type = re.ReplaceAllString(s.Content, "$2")
			s.Model = re.ReplaceAllString(s.Content, "$3")
			s.Description = re.ReplaceAllString(s.Content, "$5")
		}
	}
}

type SwaggerComment struct {
	FuncName       string
	OriginFuncName string
	FilePath       string
	Path           string
	Method         string
	PathParams     []string
	Tags           string
	Summary        string
	Description    string
	Accept         string
	Produce        string
	Params         []SwaggerCommentParam
	Responses      []SwaggerCommentResponse
}

func (s *SwaggerComment) Parse(docs []string) error {
	re := basePattern
	re2 := spacePattern
	re3 := methodPattern
	re4 := pathPattern

	s.Params = []SwaggerCommentParam{}
	s.Responses = []SwaggerCommentResponse{}
	s.PathParams = []string{}
	for _, l := range docs {
		if re.MatchString(l) {
			prefix := strings.ToLower(re.ReplaceAllString(l, "$1"))
			content := re.ReplaceAllString(l, "$2")
			switch prefix {
			case "description":
				s.Description = content
			case "tags":
				s.Tags = content
			case "summary":
				s.Summary = content
			case "accept":
				s.Accept = strings.ToLower(content)
			case "produce":
				s.Produce = strings.ToLower(content)
			case "param":
				sp := SwaggerCommentParam{Content: content}
				sp.ParseContent()
				s.Params = append(s.Params, sp)
			case "success", "failure":
				sr := SwaggerCommentResponse{
					Content: content,
					Status:  prefix,
				}
				sr.ParseContent()
				s.Responses = append(s.Responses, sr)
			case "router":
				routeString := re.ReplaceAllString(content, "$1")
				routeStrings := strings.Split(re2.ReplaceAllString(routeString, " "), " ")
				s.Path = routeStrings[0]
				s.Method = re3.ReplaceAllString(routeStrings[1], "")
				routes := strings.Split(s.Path, "/")
				for _, r := range routes {
					if re4.MatchString(r) {
						s.PathParams = append(s.PathParams, re4.ReplaceAllString(r, "$1"))
					}
				}
			}
		}
	}

	if s.Path == "" {
		return fmt.Errorf(`FATAL: invalid swagger definition. @Router not found at %s`, s.FilePath)
	}

	for _, param := range s.PathParams {
		hasParam := false
		for _, sp := range s.Params {
			if sp.In == "path" && param == sp.Name {
				hasParam = true
			}
		}

		if !hasParam {
			return fmt.Errorf(`FATAL: invalid swagger definition. @Param %s is not defined at %s`, param, s.FilePath)
		}
	}

	rex := regexp.MustCompile(`(?i)` + s.Tags)
	if !rex.MatchString(s.OriginFuncName) {
		return fmt.Errorf(`FATAL: invalid swagger definition. @Tags %s is not relevan with function %s at %s`,
			s.Tags, s.OriginFuncName, s.FilePath)
	}

	baseFile := filepath.Base(s.FilePath)
	baseMethod := filePattern.ReplaceAllString(baseFile, "$2")
	if !strings.EqualFold(baseMethod, s.Method) {
		return fmt.Errorf(`invalid swagger docs method %s. endpoint method must be %s %s`,
			baseMethod, s.Method, s.ToString())
	}

	return nil
}

func (s *SwaggerComment) ParseFile(filePath string) error {
	fset := token.NewFileSet()
	s.FilePath = filePath
	s.FuncName = toPascal(filePath)

	f, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if nil != err {
		return err
	}

	packg := &ast.Package{
		Name:  "Any",
		Files: make(map[string]*ast.File),
	}
	packg.Files[filePath] = f

	importPath, _ := filepath.Abs("/")
	dc := doc.New(packg, importPath, doc.AllMethods)
	functions := []string{}
	for _, c := range dc.Funcs {
		docs := strings.Split(c.Doc, "\n")
		s.OriginFuncName = c.Name
		if c.Name == s.FuncName && len(docs) > 0 {
			return s.Parse(docs)
		}
		functions = append(functions, c.Name)
	}

	funcs := strings.Join(functions, "\n- ")
	if len(functions) == 1 {
		f0 := functions[0]
		funcs = funcs + "\nhints: \n"
		funcs = funcs + fmt.Sprintf(`- rename function %s to %s`, f0, s.FuncName)
	}

	return fmt.Errorf("\nWARNING:\nfile %s\nhas no function with name %s.\navailable functions: \n- %s",
		s.FilePath,
		s.FuncName,
		funcs)
}

func (s *SwaggerComment) ToString() string {
	j, _ := json.MarshalIndent(s, "", "  ")
	return string(j)
}

func TestController(t *testing.T) {
	files := getAllGoFiles(".")
	fatals := []string{}
	for f := range files {
		sw := SwaggerComment{}
		if err := sw.ParseFile(files[f]); nil != err {
			if strings.HasPrefix(err.Error(), "FATAL:") {
				fatals = append(fatals, err.Error())
			} else {
				fmt.Println(err)
			}
		}
	}

	if len(fatals) > 0 {
		fmt.Println("SOME CODE NEED TO BE FIXED:")
		fmt.Println(strings.Join(fatals, "\n"))
		if os.Getenv("FAIL_ON_CONTROLLER") == "1" {
			os.Exit(1)
		}
	}
}
