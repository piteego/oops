package example

import (
	"github.com/piteego/oops"
	"github.com/piteego/oops/kind"
	"github.com/piteego/oops/severity"
)

type Metadata struct {
	oops.Metadata
	Id    string
	Retry bool
}

type MetaStandard struct {
	oops.Metadata
	kind.Code
	severity.Level
	Retry bool
}
