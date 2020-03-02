package main_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"

	"go-restapi/routes"
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

var httpCases = []HttpTestCase{

	{
		method:               "POST",
		path:                 "http://localhost:8080/:user/register",
		expectedStatus:       http.StatusOK,
		expectedResponseBody: "../mockdata/register200.json",
	},

	{
		method:               "POST",
		path:                 "http://localhost:8080/user/register",
		expectedStatus:       http.StatusBadRequest,
		expectedResponseBody: "../mockdata/register400.json",
	},

	{
		method:               "POST",
		path:                 "http://localhost:8080/user/register",
		expectedStatus:       http.StatusUnprocessableEntity,
		expectedResponseBody: "../mockdata/register422.json",
	},

	{
		method:               "POST",
		path:                 "user/find",
		expectedStatus:       http.StatusOK,
		expectedResponseBody: "../mockdata/find200.json",
	},

	{
		method:               "POST",
		path:                 "http://localhost:8080/user/find",
		expectedStatus:       http.StatusBadRequest,
		expectedResponseBody: "../mockdata/find400.json",
	},

	{
		method:               "POST",
		path:                 "http://localhost:8080/user/find",
		expectedStatus:       http.StatusUnprocessableEntity,
		expectedResponseBody: "../mockdata/find422.json",
	},

	{
		method:               "POST",
		path:                 "http://localhost:8080/user/find2",
		expectedStatus:       http.StatusOK,
		expectedResponseBody: "../mockdata/find200.json",
	},

	{
		method:               "POST",
		path:                 "http://localhost:8080/user/find2",
		expectedStatus:       http.StatusBadRequest,
		expectedResponseBody: "../mockdata/find400.json",
	},

	{
		method:               "POST",
		path:                 "http://localhost:8080/user/find2",
		expectedStatus:       http.StatusUnprocessableEntity,
		expectedResponseBody: "../mockdata/find422.json",
	},
}

func NotTestRestAPI(t *testing.T) {
	t.Log("\n\n\n\n\n\n\n\n\n\n\n\n###########  ################ ############### TestRESTAPI()")

	globalT = t
	for _, testCase := range httpCases {

		t.Log("\n\n\n\n\n\n###########  NEW REQUEST ############### ")
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

func performRequest(r http.Handler, req *http.Request, statusCode int, expectedJSON string) *httptest.ResponseRecorder {
	writer := httptest.NewRecorder()
	writer.WriteHeader(statusCode)

	globalT.Log("PerformRequst:", expectedJSON)
	body := handlerBody(expectedJSON)
	r.ServeHTTP(writer, req)
	io.WriteString(writer, body)
	return writer
}

func TestAPI(t *testing.T) {

	t.Log("\n\n\n\n\n\n###########  ################ ###############  TestAPI")
	globalT = t
	router := routes.SetupRouter()

	for _, testCase := range httpCases {

		t.Log("\n\n\n###########  NEW REQUEST ############### ")
		t.Log("ResponseBody: Method: ", testCase.method, " Path: ", testCase.path, " StatusCode: ", testCase.expectedStatus)

		request := httptest.NewRequest(testCase.method, testCase.path, nil)
		writer := performRequest(router, request, testCase.expectedStatus, testCase.expectedResponseBody)
		response := writer.Result()
		body, _ := ioutil.ReadAll(response.Body)
		t.Log("ResponseBody: ", string(body))

		expectedJSON := handlerBody(testCase.expectedResponseBody)
		assert.Equal(t, expectedJSON, string(body))
		//var response map[string]string
		//json.Unmarshal([]byte(writer.Result()))

	}
}

/*




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


*/
