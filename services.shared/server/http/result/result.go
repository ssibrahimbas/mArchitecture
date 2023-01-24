package result

type Result struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type DetailResult struct {
	Status int `json:"status"`
	Detail any `json:"detail"`
	Result
}

func (r *Result) Error() string {
	return r.Message
}

func (r *DetailResult) Error() string {
	return r.Message
}

func Success(m string, c ...int) *Result {
	code := 200
	if len(c) > 0 {
		code = c[0]
	}
	return &Result{
		Message: m,
		Status:  code,
	}
}

func Error(m string, c ...int) *Result {
	code := 400
	if len(c) > 0 {
		code = c[0]
	}
	return &Result{
		Message: m,
		Status:  code,
	}
}

func SuccessDetail(m string, d any, c ...int) *DetailResult {
	code := 200
	if len(c) > 0 {
		code = c[0]
	}
	return &DetailResult{
		Detail: d,
		Result: Result{Message: m, Status: code},
	}
}

func ErrorDetail(m string, d any, c ...int) *DetailResult {
	code := 400
	if len(c) > 0 {
		code = c[0]
	}
	return &DetailResult{
		Detail: d,
		Result: Result{Message: m,Status: code},
	}
}
