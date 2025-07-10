package example

import "errors"

var (
	Error                 = errors.New("an error")
	ErrorCause            = errors.New("a cause error")
	BuiltinErr            = errors.New("a builtin error")
	OsErrNotExist         = errors.New("file does not exist")
	RedisCacheMissed      = errors.New("redis cache missed")
	GormErrDuplicatedKey  = errors.New("gorm duplicated key")
	GormErrRecordNotFound = errors.New("gorm record not found")
)
