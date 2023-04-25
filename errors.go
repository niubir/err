package errors

import (
	"fmt"
	"io"
)

type Error interface {
	Error() string
	Format(s fmt.State, verb rune)
	Code() string
	Cause() error
	Stack() string
}

type codeError struct {
	code  string
	cause error
	stack *stack
}

func (e *codeError) Error() string {
	var s string
	if e.code != "" {
		s += "code " + e.code + ": "
	}
	s += e.cause.Error()
	return s
}

func (e *codeError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "code %s: %+v", e.code, e.Cause())
			e.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, e.Error())
	case 'q':
		fmt.Fprintf(s, "%q", e.Error())
	}
}

func (e *codeError) Code() string {
	return e.code
}

func (e *codeError) Cause() error {
	return e.cause
}

func (e *codeError) Stack() string {
	var stack string
	for _, pc := range *e.stack {
		f := Frame(pc)
		stack += fmt.Sprintf("\n%+v", f)
	}
	return stack
}

func New(code string, cause error) error {
	return &codeError{
		code:  code,
		cause: cause,
		stack: callers(),
	}
}

func Code(err error) string {
	type coder interface {
		Code() string
	}
	i, ok := err.(coder)
	if ok {
		return i.Code()
	}
	return ""
}

func Cause(err error) error {
	type causer interface {
		Cause() error
	}
	for err != nil {
		c, ok := err.(causer)
		if !ok {
			break
		}
		err = c.Cause()
	}
	return err
}

func Stack(err error) string {
	type stacker interface {
		Stack() string
	}
	i, ok := err.(stacker)
	if ok {
		return i.Stack()
	}
	return ""
}
