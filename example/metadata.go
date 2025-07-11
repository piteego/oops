package example

import (
	"github.com/piteego/oops/kind"
	"github.com/piteego/oops/severity"
	"github.com/piteego/oops/v1.0"
)

type Metadata struct {
	v1_0.Metadata
	Id    string
	Retry bool
}

type MetaKind struct {
	v1_0.Metadata
	kind.Code
	Retry bool
}

type MetaSeverity struct {
	v1_0.Metadata
	severity.Level
	Retry bool
}

type MetaKindAndSeverity struct {
	v1_0.Metadata
	kind.Code
	severity.Level
	Retry bool
}
