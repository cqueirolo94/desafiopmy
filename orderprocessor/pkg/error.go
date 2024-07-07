package pkg

import "errors"

var (
	ErrNilPointer     = errors.New("received nil pointer")
	ErrWrongOrderType = errors.New("unexpected order type")
)
