package oops

import "github.com/piteego/oops/internal"

type metadata interface{ mustBeEmbedToBeMetadata() }

func With[T metadata](metadata T) metaOption {
	return func(e error) error {
		if e == nil {
			panic("trying to set metadata to nil error")
		}
		switch err := e.(type) {
		case *internal.Error[T]:
			err.Data = metadata
			return err

		case *Error[T]:
			err.meta = metadata
			return err

		default:
			return &Error[T]{source: err, meta: metadata}
		}
	}
}

type (
	CausedBy   = internal.CauseOption
	Standard   = internal.AllOptions
	Metadata   struct{}
	metaOption func(error) error
)

func (Metadata) mustBeEmbedToBeMetadata() {}
