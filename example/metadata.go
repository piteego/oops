package example

import (
	"github.com/piteego/oops"
	"github.com/piteego/oops/kind"
)

type Metadata struct {
	oops.Metadata
	Id    string
	Retry bool
}

type MetadataAndCode struct {
	oops.Metadata
	Code  kind.Code
	Retry bool
}

func (m MetadataAndCode) Kind() (kind.Code, bool) { return m.Code, true }
