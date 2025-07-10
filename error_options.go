package oops

import (
	"github.com/piteego/oops/internal"
	"github.com/piteego/oops/kind"
	"github.com/piteego/oops/severity"
)

type (
	metadata   interface{ mustBeEmbedToBeMetadata() }
	Metadata   struct{}
	metaOption func(error) error
)

func With[T metadata](metadata T) metaOption {
	return func(e error) error {
		if e == nil {
			panic("trying to set metadata to nil error")
		}
		switch err := e.(type) {
		case *internal.Error[T]:
			err.Data = metadata
			return err

		case *metaError[T]:
			err.meta = metadata
			return err

		default:
			return &metaError[T]{source: err, meta: metadata}
		}
	}
}

func Standard(code kind.Code, level severity.Level, cause error) internal.AllOptions {
	return internal.AllOptions{
		Code:  code,
		Level: level,
		Cause: cause,
	}
}

type (
	CausedBy = internal.CauseOption
)

func (Metadata) mustBeEmbedToBeMetadata() {}
