package oops

func New[M metadata](msg string, data M) error {
	return &richError[M]{msg: msg, meta: data}
}

type (
	metadata interface{ errorMetadata() }
	Metadata struct{}
)

func (Metadata) errorMetadata() {}

type richError[T any] struct {
	msg  string
	meta T
}

func (err *richError[_]) Error() string { return err.msg }
func (err *richError[_]) Kind(sanitize bool) Kind {
	switch meta := any(err.meta).(type) {
	case kindReader:
		return meta.Kind(sanitize)

	default:
		return UnknownKind
	}
}
func (err *richError[T]) Severity(sanitize bool) Level {
	switch meta := any(err.meta).(type) {
	case levelReader:
		return meta.Severity(sanitize)

	default:
		return UnknownSeverityLevel
	}
}
func (err *richError[T]) Cause() error {
	switch meta := any(err.meta).(type) {
	case Cause:
		return meta.err

	case cause:
		return meta.Cause()

	default:
		return nil
	}
}
