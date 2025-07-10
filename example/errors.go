package example

import "errors"

var (
	ErrorCause            = errors.New("a cause error")
	OsErrNotExist         = errors.New("file does not exist")
	RedisCacheMissed      = errors.New("redis cache missed")
	GormErrDuplicatedKey  = errors.New("gorm duplicated key")
	GormErrRecordNotFound = errors.New("gorm record not found")
)
