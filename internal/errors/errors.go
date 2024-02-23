package errors

import "errors"

var (
	ErrSomethingWentWrong = errors.New("something went wrong")
	ErrWrongLogin         = errors.New("username or password is wrong")
)
