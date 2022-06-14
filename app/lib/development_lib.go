//go:build !production
// +build !production

package lib

/**
 * NOTICE
 *
 * Feel free to create your own function here for Unit Testing purpose
 * Also make sure you provide unit tests of the functions you create
 * Do not use any of the functions described here for production
 */
import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
)

// HTTPRequest create a simple request object
func HTTPRequest(method string, path string, headers map[string]string, body ...string) *http.Request {
	var payload io.Reader
	if len(body) == 1 && body[0] != "" {
		payload = bytes.NewReader([]byte(body[0]))
	}
	request := httptest.NewRequest(method, path, payload)
	if nil != headers {
		for i := range headers {
			request.Header.Add(i, headers[i])
		}
	} else {
		request.Header.Add("Accept", "application/json")
		request.Header.Add("Content-type", "application/json")
	}

	return request
}

// GetTest get response for GetTest
func GetTest(app *fiber.App, path string, headers map[string]string) (*http.Response, map[string]interface{}, error) {
	var result map[string]interface{}
	response, err := app.Test(HTTPRequest("GET", path, headers), 500)
	defer func() {
		response.Body.Close()
	}()

	if nil == err {
		if bte, err := ioutil.ReadAll(response.Body); nil == err {
			json.Unmarshal(bte, &result)
		}
	}

	return response, result, err
}

// PostTest get response for PostTest
func PostTest(app *fiber.App, path string, headers map[string]string, body ...string) (*http.Response, map[string]interface{}, error) {
	var result map[string]interface{}
	response, err := app.Test(HTTPRequest("POST", path, headers, body...), 500)
	defer func() {
		response.Body.Close()
	}()

	if nil == err {
		if bte, err := ioutil.ReadAll(response.Body); nil == err {
			json.Unmarshal(bte, &result)
		}
	}

	return response, result, err
}

// PutTest get response for PutTest
func PutTest(app *fiber.App, path string, headers map[string]string, body ...string) (*http.Response, map[string]interface{}, error) {
	var result map[string]interface{}
	response, err := app.Test(HTTPRequest("PUT", path, headers, body...), 500)
	defer func() {
		response.Body.Close()
	}()

	if nil == err {
		if bte, err := ioutil.ReadAll(response.Body); nil == err {
			json.Unmarshal(bte, &result)
		}
	}

	return response, result, err
}

// DeleteTest get response for DeleteTest
func DeleteTest(app *fiber.App, path string, headers map[string]string) (*http.Response, map[string]interface{}, error) {
	var result map[string]interface{}
	response, err := app.Test(HTTPRequest("DELETE", path, headers), 500)
	defer func() {
		response.Body.Close()
	}()

	if nil == err {
		if bte, err := ioutil.ReadAll(response.Body); nil == err {
			json.Unmarshal(bte, &result)
		}
	}

	return response, result, err
}
