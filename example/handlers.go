package example

import (
	"errors"
	"github.com/piteego/oops/v0"
)

func HandleRepoErr(entity string) v0.Handler {
	return func(err error) *v0.Error {
		if err == nil {
			return nil
		}
		if errors.Is(err, RedisCacheMissed) {
			return v0.New(entity+" not found", v0.Tag(NotFound.Error)).(*v0.Error)
		}
		if errors.Is(err, GormErrDuplicatedKey) {
			return v0.New("duplicated "+entity, v0.Tag(Duplication.Error)).(*v0.Error)
		}
		if errors.Is(err, GormErrRecordNotFound) {
			return v0.New(entity+" not found", v0.Tag(NotFound.Error)).(*v0.Error)
		}
		return v0.New("something went wrong", v0.Tag(Internal.Error)).(*v0.Error)
	}
}
