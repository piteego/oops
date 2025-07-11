package example

import (
	"github.com/piteego/oops"
)

type Metadata struct {
	oops.Metadata
	Id    string
	Retry bool
}

type MetaKind struct {
	oops.Metadata
	oops.Kind
	Retry bool
}

type MetaSeverity struct {
	oops.Metadata
	oops.Level
	Retry bool
}

type MetaKindAndSeverity struct {
	oops.Metadata
	oops.Kind
	oops.Level
	Retry bool
}
