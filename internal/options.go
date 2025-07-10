package internal

import (
	"github.com/piteego/oops/kind"
	"github.com/piteego/oops/severity"
)

type (
	Option interface {
		kind.Code | severity.Level | Cause | AllOptions
	}
	Cause      struct{ Error error }
	AllOptions struct {
		Code  kind.Code
		Level severity.Level
		Cause error
	}
)

func (opt AllOptions) Kind() kind.Code          { return opt.Code }
func (opt AllOptions) Severity() severity.Level { return opt.Level }
