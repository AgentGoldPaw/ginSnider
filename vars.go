package gin_unit_test

import (
	"log"
	"net/http"
)

var (
	router        http.Handler
	logging       *log.Logger
	globalHeaders map[string]string
	DefaultLogger bool
)
