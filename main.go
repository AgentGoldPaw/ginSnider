package gin_unit_test

import (
	"encoding/json"
	"github.com/golden-protocol/gin_unit_test/utils"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
)

func SetRouter(r http.Handler) {
	router = r
}

func SetLog(l *log.Logger) {
	logging = l
}

// printf log
func printfLog(format string, v ...interface{}) {
	if DefaultLogger && logging == nil {
		logging = log.Default()
	}
	if logging == nil {
		return
	}
	logging.Printf(format, v...)
}

// invoke handler
func invokeHandler(req *http.Request) (writer *http.Response, bodyByte []byte, err error) {

	// initialize response record
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	// extract the response from the response record
	result := w.Result()
	defer result.Body.Close()

	// extract response body
	bodyByte, err = ioutil.ReadAll(result.Body)

	return w.Result(), bodyByte, err
}

func TestOrdinaryHandler(testRequest *TestOrdinaryHandlerStruct) (resp *http.Response, bodyByte []byte, err error) {
	if router == nil {
		err = ErrRouterNotSet
		return
	}

	printfLog("TestOrdinaryHandler\tRequest:\t%v:%v,\trequestBody:%v\n", testRequest.Method, testRequest.Api, testRequest.Param)

	// make request
	req, err := utils.MakeRequest(testRequest.Method, testRequest.Mime, testRequest.Api, testRequest.Param)
	if err != nil {
		return
	}
	// add the customed headers
	if testRequest.Headers != nil && len(testRequest.Headers) > 0 {
		addHeadersToRequest(req.Header, mergeHeaders(testRequest.Headers))
	} else {
		addHeadersToRequest(req.Header, globalHeaders)
	}

	if !testRequest.useAuth {
		// merge headers call above merges the override headers into global
		// we only need to pass global here
		skipAuthHeaders(req.Header, globalHeaders)
	}

	// invoke handler
	resp, bodyByte, err = invokeHandler(req)

	printfLog("TestOrdinaryHandler\tResponse:\t%v:%v\tResponse:%v\n\n\n", testRequest.Method, testRequest.Api, string(bodyByte))
	return
}

func TestHandlerUnMarshalResp(testRequest *TestOrdinaryHandlerStruct) (int, error) {
	resp, bodyByte, err := TestOrdinaryHandler(testRequest)
	if err != nil {
		return 0, err
	}
	err = json.Unmarshal(bodyByte, testRequest.Response)
	return resp.StatusCode, err
}
