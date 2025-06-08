package example

import (
	"fmt"
	"github.com/piteego/oops"
)

type Metadata struct {
	oops.Metadata
	Code  int
	Retry bool
}

func (m Metadata) String() string {
	return fmt.Sprintf("{Code: %d, Retry: %t}", m.Code, m.Retry)
}
