package models
type Result struct {
	Code int
	Msg  string
	Data interface{}
}
func NewResult(c int, m string, dt interface{}) *Result {
	r := &Result{Code: c, Data: dt}

	if e, ok := dt.(error); ok {
		if m == "" {
			r.Msg = e.Error()
		}
	} else {
		r.Msg = "SUCCESS"
	}

	return r
}