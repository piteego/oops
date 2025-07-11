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

type MetaKind struct {
	oops.Metadata
	kind.Code
	Retry bool
}

type MetaSeverity struct {
	oops.Metadata
	severity.Level
	Retry bool
}

type MetaKindAndSeverity struct {
	oops.Metadata
	kind.Code
	severity.Level
	Retry bool
}
