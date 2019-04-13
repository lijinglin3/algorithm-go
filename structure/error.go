package structure

import "errors"

var (
	ErrorIndexOutOfRange       = errors.New("index out of range")
	ErrorFullArray             = errors.New("full array")
	ErrorNilPointerDereference = errors.New("nil pointer")
	ErrorNotFound              = errors.New("not found")
)
