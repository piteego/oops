package v1_3

import (
	"errors"
	"github.com/piteego/oops/internal"
	"github.com/piteego/oops/kind"
	"github.com/piteego/oops/severity"
)

type metadata interface{ errorMetadata() }

func Get[D interface {
	kind.Code | severity.Level | Cause
}](err error) D {
	if err == nil {
		var zero D
		return zero
	}
	switch e := err.(type) {
	case *wrap[D]:
		return e.data

	default:
		return Get[D](errors.Unwrap(err))
	}
}

type Metadata struct{}

func (Metadata) errorMetadata() {}

type Cause = internal.CauseOption
