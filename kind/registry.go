package kind

import "sync"

var (
	once     = new(sync.Once)
	registry map[Code]error
)

func Register(kinds map[Code]error) {
	if kinds == nil || len(kinds) == 0 {
		return
	}
	once.Do(func() { registry = kinds })
}
