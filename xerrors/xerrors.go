package xerrors

import (
	"fmt"
	"runtime"
)

// HTTP相关错误代码
const (
	// Client Errors
	ErrBadRequest = iota + 400
	ErrUnauthorized
	ErrForbidden
	ErrNotFound
	ErrMethodNotAllowed
	ErrNotAcceptable
	ErrRequestTimeout
	ErrConflict
	ErrGone
	ErrPreconditionFailed
	ErrPayloadTooLarge
	ErrUnsupportedMediaType
	ErrTooManyRequests
	// Server Errors
	ErrInternalServerError = iota + 500
	ErrNotImplemented
	ErrBadGateway
	ErrServiceUnavailable
	ErrGatewayTimeout
	ErrHTTPVersionNotSupported
	ErrVariantAlsoNegotiates
	ErrInsufficientStorage
	ErrLoopDetected
	ErrNotExtended
	ErrNetworkAuthenticationRequired
)

// 数据库相关错误代码
const (
	ErrDatabaseConnectionFailed = iota + 6001
	ErrDatabaseQueryFailed
	ErrDatabaseUpdateFailed
	ErrDatabaseNotFound
	ErrDuplicateRecord
	ErrInsufficientDatabasePermissions
	ErrDatabaseTimeout
	ErrDatabaseTransactionFailed
	ErrDatabaseIntegrityConstraintViolation
	ErrDatabaseDiskFull
	ErrDatabaseServiceUnavailable
)

// 网络相关错误代码
const (
	ErrDNSResolutionFailed = iota + 7001
	// 其他网络相关错误代码可以在这里添加
)

// Error represents a unified error type.
type Error struct {
	Code    int    // Error code
	Message string // Error message
	Cause   error  // Original error causing this error
	Stack   string // Error stack trace
}

// New creates a new Error instance.
func New(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Stack:   stackTrace(),
	}
}

// Wrap wraps an existing error with additional context.
func Wrap(err error, code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Cause:   err,
		Stack:   stackTrace(),
	}
}

// Unwrap returns the next error in error chain.
func (e *Error) Unwrap() error {
	return e.Cause
}

// Error implements the error interface.
func (e *Error) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

// stackTrace returns the stack trace as a string.
func stackTrace() string {
	// For simplicity, keeping the same stack trace function
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	return string(buf[:n])
}
