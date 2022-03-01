package gin_unit_test

type TestOrdinaryHandlerStruct struct {
	Method   string
	Api      string
	Mime     string
	Param    interface{}
	Headers  map[string]string
	useAuth  bool
	Response interface{}
}

func (t *TestOrdinaryHandlerStruct) SetAuth(auth bool) {
	t.useAuth = auth
}
