package example

import (
	"github.com/piteego/oops"
)

var ErrMap = oops.Map{
	RedisCacheMissed:      oops.New("cache key not found", oops.Tag(NotFound.Error)).(*oops.Error),
	GormErrDuplicatedKey:  oops.New("duplicated entity", oops.Tag(Duplication.Error)).(*oops.Error),
	GormErrRecordNotFound: oops.New("entity not found", oops.Tag(NotFound.Error)).(*oops.Error),
}
