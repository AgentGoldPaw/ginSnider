package gin_unit_test

import (
	"github.com/golden-protocol/super-pancake/slices"
	"net/http"
)

func SetGlobalHeaders(headers map[string]string) {
	globalHeaders = headers
}

func AddHeader(key, value string, skip ...string) {
	if slices.IndexOfString(skip, key) > -1 {
		globalHeaders[key] = value
	} else if _, ok := globalHeaders[key]; ok {
		delete(globalHeaders, key)
	}
}

func mergeHeaders(extraHeaders map[string]string, skip ...string) map[string]string {
	for key, value := range extraHeaders {
		AddHeader(key, value, skip...)
	}
	return globalHeaders
}

func addHeadersToRequest(requestHeaders http.Header, headers map[string]string, skip ...string) {
	for key, value := range headers {
		if slices.IndexOfString(skip, key) > -1 {
			requestHeaders.Add(key, value)
		} else if requestHeaders.Get(key) != "" {
			requestHeaders.Del(key)
		}
	}
}

func skipAuthHeaders(requestHeaders http.Header, headers map[string]string) {
	addHeadersToRequest(requestHeaders, mergeHeaders(headers, "Authorization"))
}
