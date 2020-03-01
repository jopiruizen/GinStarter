package main_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

/* REST API TESTS */

type HttpTestCase struct {
	method               string
	path                 string
	jsonParams           string
	expectedStatus       int
	expectedResponseBody string
	handler              func(http.ResponseWriter, *http.Request)
}

var globalT *testing.T

func handlerBody(source string) string {
	jsonFile, _ := os.Open(source)
	body, _ := ioutil.ReadAll(jsonFile)
	globalT.Log("Load Test JSON Source", source)
	return string(body)
}

func registerTestHandlers200(writer http.ResponseWriter, request *http.Request) {

	globalT.Log("\n\n\n\n\n##################################")
	globalT.Log("registerTestHandlers200()")

	writer.WriteHeader(http.StatusOK)
	body := handlerBody("../mockdata/register200.json")
	io.WriteString(writer, string(body))
}

func registerTestHandlers400(writer http.ResponseWriter, request *http.Request) {
	globalT.Log("\n\n\n\n\n##################################")
	globalT.Log("registerTestHandlers400()")

	writer.WriteHeader(http.StatusBadRequest)
	body := handlerBody("../mockdata/register400.json")
	io.WriteString(writer, body)
}

func registerTestHandlers422(writer http.ResponseWriter, request *http.Request) {
	globalT.Log("\n\n\n\n\n##################################")
	globalT.Log("registerTestHandlers422()")

	writer.WriteHeader(http.StatusUnprocessableEntity)
	body := handlerBody("../mockdata/register422.json")
	io.WriteString(writer, body)
}

func findTestHandlers200(writer http.ResponseWriter, request *http.Request) {
	globalT.Log("\n\n\n\n\n##################################")
	globalT.Log("findTestHandlers200()")
	writer.WriteHeader(http.StatusOK)
	body := handlerBody("../mockdata/find200.json")
	io.WriteString(writer, body)
}

func findTestHandlers400(writer http.ResponseWriter, request *http.Request) {
	globalT.Log("\n\n\n\n\n##################################")
	globalT.Log("findTestHandlers400()")
	writer.WriteHeader(http.StatusBadRequest)
	body := handlerBody("../mockdata/find400.json")
	io.WriteString(writer, body)
}

func findTestHandlers422(writer http.ResponseWriter, request *http.Request) {
	globalT.Log("\n\n\n\n\n##################################")
	globalT.Log("findTestHandlers422()")
	writer.WriteHeader(http.StatusUnprocessableEntity)
	body := handlerBody("../mockdata/find422.json")
	io.WriteString(writer, body)
}

func find2TestHandlers200(writer http.ResponseWriter, request *http.Request) {
	globalT.Log("\n\n\n\n\n##################################")
	globalT.Log("find2TestHandlers200()")
	writer.WriteHeader(http.StatusOK)
	body := handlerBody("../mockdata/find200.json")
	io.WriteString(writer, body)
}

func find2TestHandlers400(writer http.ResponseWriter, request *http.Request) {
	globalT.Log("\n\n\n\n\n##################################")
	globalT.Log("find2TestHandlers400()")
	writer.WriteHeader(http.StatusBadRequest)
	body := handlerBody("../mockdata/find400.json")
	io.WriteString(writer, body)
}

func find2TestHandlers422(writer http.ResponseWriter, request *http.Request) {
	globalT.Log("\n\n\n\n\n##################################")
	globalT.Log("findTestHandlers422()")
	writer.WriteHeader(http.StatusUnprocessableEntity)
	body := handlerBody("../mockdata/find422.json")
	io.WriteString(writer, body)
}

var httpCases = []HttpTestCase{

	{
		method:               "POST",
		path:                 "http://github.com/go-restapi/user/register",
		expectedStatus:       http.StatusOK,
		expectedResponseBody: "../mockdata/register200.json",
		handler:              registerTestHandlers200,
	},

	{
		method:               "POST",
		path:                 "http://github.com/go-restapi/user/register",
		expectedStatus:       http.StatusBadRequest,
		expectedResponseBody: "../mockdata/register400.json",
		handler:              registerTestHandlers400,
	},

	{
		method:               "POST",
		path:                 "http://github.com/go-restapi/user/register",
		expectedStatus:       http.StatusUnprocessableEntity,
		expectedResponseBody: "../mockdata/register422.json",
		handler:              registerTestHandlers422,
	},

	{
		method:               "POST",
		path:                 "http://github.com/go-restapi/user/find",
		expectedStatus:       http.StatusOK,
		expectedResponseBody: "../mockdata/find200.json",
		handler:              findTestHandlers200,
	},

	{
		method:               "POST",
		path:                 "http://github.com/go-restapi/user/find",
		expectedStatus:       http.StatusBadRequest,
		expectedResponseBody: "../mockdata/find400.json",
		handler:              findTestHandlers400,
	},

	{
		method:               "POST",
		path:                 "http://github.com/go-restapi/user/find",
		expectedStatus:       http.StatusUnprocessableEntity,
		expectedResponseBody: "../mockdata/find422.json",
		handler:              findTestHandlers422,
	},

	{
		method:               "POST",
		path:                 "http://github.com/go-restapi/user/find2",
		expectedStatus:       http.StatusOK,
		expectedResponseBody: "../mockdata/find200.json",
		handler:              find2TestHandlers200,
	},

	{
		method:               "POST",
		path:                 "http://github.com/go-restapi/user/find2",
		expectedStatus:       http.StatusBadRequest,
		expectedResponseBody: "../mockdata/find400.json",
		handler:              find2TestHandlers400,
	},

	{
		method:               "POST",
		path:                 "http://github.com/go-restapi/user/find2",
		expectedStatus:       http.StatusUnprocessableEntity,
		expectedResponseBody: "../mockdata/find422.json",
		handler:              find2TestHandlers422,
	},
}

func TestRestAPI(t *testing.T) {
	t.Log("\n\n\n\n\n\n\n\n\n\n\n\n###########  ################ ############### TestRESTAPI()")

	globalT = t
	for _, testCase := range httpCases {

		t.Log("\n\n\n\n\n\n###########  ################ ############### ")
		request := httptest.NewRequest(testCase.method, testCase.path, nil)
		writer := httptest.NewRecorder()
		testCase.handler(writer, request)
		response := writer.Result()

		body, _ := ioutil.ReadAll(response.Body)
		t.Log(fmt.Sprintf("%+v\n", response))

		if testCase.expectedStatus != response.StatusCode {
			t.Fatal("Status Code Failed at ", testCase.path, " with Code: ", testCase.expectedStatus)
		}

		expectedJSON := handlerBody(testCase.expectedResponseBody)
		t.Log("\n\n#####\n - compare?")
		t.Log("ResponseBody: ", string(body))
		t.Log(" Expected: ", expectedJSON)
		if string(body) != expectedJSON {
			t.Fatal("Response Failed on ", testCase.path, " with Status Code: ", testCase.expectedStatus)
		}
	}
}
