package internal

type Error[T any] struct {
	Msg  string
	Data T
}

func (err *Error[_]) Error() string { return err.Msg }
