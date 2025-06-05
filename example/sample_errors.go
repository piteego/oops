package example

import "errors"

var (
	OsErrNotExist         = errors.New("file does not exist")
	RedisCacheMissed      = errors.New("redis cache missed")
	GormErrDuplicatedKey  = errors.New("gorm duplicated key")
	GormErrRecordNotFound = errors.New("gorm record not found")
)
