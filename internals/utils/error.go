package utils

type CustError struct {
	ErrMsg     ErrorMsg
	StatusCode int
}

type ErrorMsg struct {
	ErrMsg string
}

func NewCustError(err error, statusCode int) *CustError {
	return &CustError{ErrMsg: ErrorMsg{ErrMsg: err.Error()}, StatusCode: statusCode}
}
