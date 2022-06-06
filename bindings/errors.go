package bindings

import (
	"errors"
	"strings"
)

var (
	ErrUsernameEmpty = errors.New("username cannot be empty")
	ErrPasswordEmpty = errors.New("password cannot be empty")
)

type RequestErrors struct {
	errs []error
}

func (re *RequestErrors) Append(err error) {
	re.errs = append(re.errs, err)
}

func (re *RequestErrors) Len() int {
	return len(re.errs)
}

func (re *RequestErrors) Error() string {
	errStr := []string{}
	for _, e := range re.errs {
		errStr = append(errStr, e.Error())
	}
	return strings.Join(errStr, ", ")
}
