package internal

import (
	"github.com/piteego/oops/kind"
	"github.com/piteego/oops/severity"
)

type (
	Option interface {
		kind.Code | severity.Level | CauseOption | AllOptions
	}
	CauseOption struct{ Error error }
	AllOptions  struct {
		Code  kind.Code
		Level severity.Level
		Cause error
	}
)
