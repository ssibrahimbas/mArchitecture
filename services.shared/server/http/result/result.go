package result

type Result struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
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

func Success(m string, c int) *Result {
	return &Result{
		Message: m,
		Code:    c,
	}
}

func Error(m string, c int) *Result {
	return &Result{
		Message: m,
		Code:    c,
	}
}

func SuccessDetail(m string, d any, c int) *DetailResult {
	return &DetailResult{
		Detail: d,
		Result: Result{Message: m, Code: c},
	}
}

func ErrorDetail(m string, d any, c int) *DetailResult {
	return &DetailResult{
		Detail: d,
		Result: Result{Message: m, Code: c},
	}
}
