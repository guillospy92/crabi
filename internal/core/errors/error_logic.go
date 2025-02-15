package errorlogic

import (
	"net/http"

	"github.com/pkg/errors"
)

const (
	// LoginErrorNotFound login not found
	LoginErrorNotFound = "LoginErrorNotFound"

	// UserErrorBlocked  user is blocked
	UserErrorBlocked = "UserErrorBlocked"

	// UserErrorExists user error exists persistence
	UserErrorExists = "UserErrorExists"

	// TokenExpired token expired
	TokenExpired = "TokenExpired"
)

var errorByCode = map[int]map[string]string{
	http.StatusNotFound: {
		"message":    "resource not found",
		"error_code": "ResourceNotFoundError",
	},
	http.StatusInternalServerError: {
		"message":    "unexpected error",
		"error_code": "UnexpectedError",
	},
}

// ResponseErrorLogic responsé error struct json
type ResponseErrorLogic struct {
	Message    string `json:"message"`
	ErrorCode  string `json:"error_code"`
	StatusCode int    `json:"status_code"`
}

// ErrorLogic error of our domain that can operate on the entire application
type ErrorLogic struct {
	Message    string
	Err        error
	ErrCode    string
	StatusCode int
}

// Error implements interface error
func (e *ErrorLogic) Error() string {
	return e.Message
}

// GenerateErrorJSON get ErrorLogic and convert attributes in response of json
func (e *ErrorLogic) GenerateErrorJSON() *ResponseErrorLogic {
	return &ResponseErrorLogic{
		Message:    e.Message,
		ErrorCode:  e.ErrCode,
		StatusCode: e.StatusCode,
	}
}

func (e *ErrorLogic) searchCodeError() {
	if e.ErrCode != "" && e.Message != "" {
		return
	}

	errByCode, ok := errorByCode[e.StatusCode]
	if !ok {
		errByCode = errorByCode[http.StatusInternalServerError]
		e.StatusCode = http.StatusInternalServerError
	}

	if e.ErrCode == "" {
		e.ErrCode = errByCode["error_code"]
	}

	if e.Message == "" {
		e.Message = errByCode["message"]
	}
}

// GenerateErrorSinceMessage create new instance ErrorLogic, Turning the message into an error
func GenerateErrorSinceMessage(options ...OptionalErrorGlobal) *ErrorLogic {
	errorGlobal := &ErrorLogic{}

	for _, opt := range options {
		opt(errorGlobal)
	}

	errorGlobal.Err = errors.New(errorGlobal.Message)
	errorGlobal.searchCodeError()

	return errorGlobal
}
