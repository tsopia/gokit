package xerrors

import (
	"errors"
	"testing"
)

func TestNewError(t *testing.T) {
	// 测试 New 函数是否正确创建新的 Error 实例
	err := New(404, "Resource not found")

	// 检查错误实例的属性
	if err.Code != 404 {
		t.Errorf("Expected error code 404, got %d", err.Code)
	}
	if err.Message != "Resource not found" {
		t.Errorf("Expected error message 'Resource not found', got '%s'", err.Message)
	}
	if err.Cause != nil {
		t.Error("Expected cause to be nil")
	}
	if len(err.Stack) == 0 {
		t.Error("Expected non-empty stack trace")
	}
}

func TestWrapError(t *testing.T) {
	// 模拟一个原始错误
	originalError := errors.New("database connection failed")

	// 使用 Wrap 函数将原始错误包装成新的 Error 实例
	err := Wrap(originalError, 6001, "Database connection failed")

	// 检查错误实例的属性
	if err.Code != 6001 {
		t.Errorf("Expected error code 6001, got %d", err.Code)
	}
	if err.Message != "Database connection failed" {
		t.Errorf("Expected error message 'Database connection failed', got '%s'", err.Message)
	}
	if err.Cause != originalError {
		t.Error("Expected cause to be original error")
	}
	if len(err.Stack) == 0 {
		t.Error("Expected non-empty stack trace")
	}
}

func TestErrorUnwrap(t *testing.T) {
	// 模拟一个原始错误
	originalError := errors.New("database query failed")

	// 使用 Wrap 函数将原始错误包装成新的 Error 实例
	err := Wrap(originalError, 6002, "Database query failed")

	// 使用 Unwrap 方法检查原始错误是否正确返回
	if unwrapped := errors.Unwrap(err); unwrapped != originalError {
		t.Error("Expected Unwrap to return original error")
	}
}

func TestErrorMessage(t *testing.T) {
	// 测试 Error 方法是否正确返回错误消息
	err := New(404, "Resource not found")
	expectedMessage := "[404] Resource not found"
	if errMsg := err.Error(); errMsg != expectedMessage {
		t.Errorf("Expected error message '%s', got '%s'", expectedMessage, errMsg)
	}
}

func TestStackTrace(t *testing.T) {
	// 测试 stackTrace 函数是否能够正确返回堆栈信息
	stack := stackTrace()
	if len(stack) == 0 {
		t.Error("Expected non-empty stack trace")
	}
}
func TestStackTraceOutput(t *testing.T) {
	// 测试堆栈输出是否包含正确的信息
	err := New(404, "Resource not found")
	expectedMessage := "[404] Resource not found"

	if errMsg := err.Error(); errMsg != expectedMessage {
		t.Errorf("Expected error message '%s', got '%s'", expectedMessage, errMsg)
	}
	t.Logf("%v", err.Stack)
	if stack := err.Stack; len(stack) == 0 {
		t.Error("Expected non-empty stack trace")
	}
}
