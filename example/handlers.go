package example

import (
	"errors"
	"github.com/piteego/oops"
)

func HandleRepoErr(entity string) oops.Handler {
	return func(err error) *oops.Error {
		if err == nil {
			return nil
		}
		if errors.Is(err, RedisCacheMissed) {
			return oops.New(entity+" not found", oops.Tag(NotFound.Error)).(*oops.Error)
		}
		if errors.Is(err, GormErrDuplicatedKey) {
			return oops.New("duplicated "+entity, oops.Tag(Duplication.Error)).(*oops.Error)
		}
		if errors.Is(err, GormErrRecordNotFound) {
			return oops.New(entity+" not found", oops.Tag(NotFound.Error)).(*oops.Error)
		}
		return oops.New("something went wrong", oops.Tag(Internal.Error)).(*oops.Error)
	}
}
