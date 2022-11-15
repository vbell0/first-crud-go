package resterr

import (
	"net/http"
)

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes,omitempty"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestErr) Errror() string {
	return r.Message
}

func NewRestErr(message string, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidateError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternoServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "interno_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewInternoServerValidateError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "interno_server_error",
		Code:    http.StatusInternalServerError,
		Causes:  causes,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_faund",
		Code:    http.StatusNotFound,
	}
}

func NewNotFoundValidadeError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_faund",
		Code:    http.StatusNotFound,
		Causes:  causes,
	}
}

func NewUnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

func NewUnauthorizedValidadeError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
		Causes:  causes,
	}
}

func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}

func NewForbiddenValidateError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
		Causes:  causes,
	}
}
