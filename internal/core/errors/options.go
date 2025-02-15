package errorlogic

// OptionalErrorGlobal contract for helps function ErrorLogic
type OptionalErrorGlobal func(e *ErrorLogic)

// Message set attribute Message to ErrorLogic
func Message(message string) OptionalErrorGlobal {
	return func(e *ErrorLogic) {
		e.Message = message
	}
}

// StatusCode set attribute StatusCode to ErrorLogic
func StatusCode(statusCode int) OptionalErrorGlobal {
	return func(e *ErrorLogic) {
		e.StatusCode = statusCode
	}
}

// ErrorCode set attribute ErrCode to ErrorLogic
func ErrorCode(errCode string) OptionalErrorGlobal {
	return func(e *ErrorLogic) {
		e.ErrCode = errCode
	}
}
