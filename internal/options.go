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

func (std AllOptions) GetCode() kind.Code       { return std.Code }
func (std AllOptions) GetCause() error          { return std.Cause }
func (std AllOptions) GetLevel() severity.Level { return std.Level }
