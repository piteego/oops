package example

import (
	"fmt"
	"github.com/piteego/oops/v1_0"
)

type Metadata struct {
	v1_0.Metadata
	Code  int
	Retry bool
}

func (m Metadata) String() string {
	return fmt.Sprintf("{Code: %d, Retry: %t}", m.Code, m.Retry)
}
