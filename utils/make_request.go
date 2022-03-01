package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golden-protocol/gin_unit_test/methods"
	"github.com/golden-protocol/gin_unit_test/mime"
	"io"
	"net/http"
	"strings"
)

var (
	ErrMethodNotSupported = errors.New("method is not supported")
	ErrMIMENotSupported   = errors.New("mime is not supported")
)

func MakeRequest(method, mimeType, api string, param interface{}) (request *http.Request, err error) {
	method = strings.ToUpper(method)
	mimeType = strings.ToLower(mimeType)

	switch mimeType {
	case mime.JSON:
		var (
			contentBuffer *bytes.Buffer
			jsonBytes     []byte
		)
		jsonBytes, err = json.Marshal(param)
		if err != nil {
			return
		}
		contentBuffer = bytes.NewBuffer(jsonBytes)
		request, err = http.NewRequest(method, api, contentBuffer)
		if err != nil {
			return
		}
		request.Header.Set("Content-Type", "application/json;charset=utf-8")
	case mime.FORM:
		queryStr := MakeQueryStrFrom(param)
		var buffer io.Reader

		if (method == methods.DELETE || method == methods.GET) && queryStr != "" {
			api += "?" + queryStr
		} else {
			buffer = bytes.NewReader([]byte(queryStr))
		}

		request, err = http.NewRequest(method, api, buffer)
		if err != nil {
			return
		}
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	default:
		err = ErrMIMENotSupported
		return
	}
	return
}
