package oops

type (
	cause interface{ Cause() error }
	Cause struct{ err error }
)

func (c Cause) Cause() error   { return c.err }
func (c Cause) errorMetadata() {}

func CausedBy(err error) Cause {
	if err == nil {
		panic("trying to set nil cause")
	}
	return Cause{err: err}
}

func CauseOf(err error) error {
	if err == nil {
		return nil
	}
	switch e := err.(type) {
	case cause:
		return e.Cause()

	case interface{ Unwrap() error }:
		return CauseOf(e.Unwrap())

	default:
		return nil
	}
}
