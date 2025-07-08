package example

import "github.com/piteego/oops"

type Metadata struct {
	oops.Metadata
	Id    string
	Retry bool
}
