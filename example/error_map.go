package example

import (
	"github.com/piteego/oops/v0"
)

var ErrMap = v0.Map{
	RedisCacheMissed:      v0.New("cache key not found", v0.Tag(NotFound.Error)).(*v0.Error),
	GormErrDuplicatedKey:  v0.New("duplicated entity", v0.Tag(Duplication.Error)).(*v0.Error),
	GormErrRecordNotFound: v0.New("entity not found", v0.Tag(NotFound.Error)).(*v0.Error),
}
