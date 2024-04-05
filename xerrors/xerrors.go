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

// 定义资源操作过程中可能遇到的错误类型
const (
	ErrResourceNotFound                                        = iota + 8001 // 资源未找到
	ErrResourceAlreadyExists                                                 // 资源已存在
	ErrResourceInvalid                                                       // 资源无效
	ErrResourceNotAuthorized                                                 // 无权限访问资源
	ErrResourceNotAllowed                                                    // 资源不允许操作
	ErrResourceNotImplemented                                                // 资源未实现
	ErrResourceNotSupported                                                  // 资源不被支持
	ErrResourceNotAvailable                                                  // 资源不可用
	ErrResourceNotAvailableInThisRegion                                      // 该区域中资源不可用
	ErrResourceNotAvailableInThisZone                                        // 该区域中资源不可用
	ErrResourceNotAvailableInThisEnvironment                                 // 该环境中资源不可用
	ErrResourceNotAvailableInThisEnvironmentType                             // 该环境类型中资源不可用
	ErrResourceNotAvailableInThisEnvironmentVersion                          // 该环境版本中资源不可用
	ErrResourceNotAvailableInThisEnvironmentVersionType                      // 该环境版本类型中资源不可用
	ErrResourceNotAvailableInThisEnvironmentVersionTypeVersion               // 该环境版本类型版本中资源不可用
)

// 文件不存在的err
const ErrFileNotFound = 4001

// EnvBindingErrorCode 定义绑定环境变量失败的错误码。
const (
	ErrEnvBinding = iota + 5001
	ErrUnmarshal
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
