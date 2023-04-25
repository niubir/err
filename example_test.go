package errors_test

import (
	stderrors "errors"
	"fmt"
	"testing"

	"github.com/niubir/errors"
)

var (
	example_error = errors.New("E001", stderrors.New("args error"))
)

func ExampleNew() {
	fmt.Println(example_error)

	// Output: code E001: args error
}

func ExampleError() {
	fmt.Println(example_error.Error())

	// Output: code E001: args error
}

func ExampleCode() {
	fmt.Println(errors.Code(example_error))

	// Output: E001
}

func ExampleCause() {
	fmt.Println(errors.Cause(example_error).Error())

	// Output: args error
}

func TestStack(t *testing.T) {
	t.Log(errors.Stack(example_error))
}
