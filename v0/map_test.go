package v0_test

import (
	"errors"
	"github.com/piteego/oops/example"
	"github.com/piteego/oops/v0"
	"testing"
)

func TestMap_Handle(t *testing.T) {
	var errMap = v0.Map{
		example.OsErrNotExist:         v0.New("no such file or directory", v0.Tag(example.NotFound.Error)).(*v0.Error),
		example.RedisCacheMissed:      v0.New("cache key not found", v0.Tag(example.NotFound.Error)).(*v0.Error),
		example.GormErrRecordNotFound: v0.New("entity not found", v0.Tag(example.NotFound.Error)).(*v0.Error),
	}
	for err := range errMap {
		t.Run(err.Error(), func(t *testing.T) {
			got := errMap.Handle(err)
			if !errors.Is(got, err) {
				t.Errorf("expected true, got false")
			}
			if !errors.Is(got, example.NotFound.Error) {
				t.Errorf("expected true, got false")
			}
		})
	}
	t.Run("duplicated map key", func(t *testing.T) {
		errMap[example.OsErrNotExist] = v0.New(
			"no such file or directory__duplicated!",
			v0.Tag(example.Internal.Error),
		).(*v0.Error)
		t.Logf("error map with duplicated example.OsErrNotExist key: %q", errMap[example.OsErrNotExist])
		got := errMap.Handle(example.OsErrNotExist)
		if !errors.Is(got, example.OsErrNotExist) {
			t.Errorf("expected true, got false")
		}
		if errors.Is(got, example.NotFound.Error) {
			t.Errorf("expected false, got true")
		}
		if !errors.Is(got, example.Internal.Error) {
			t.Errorf("expected true, got false")
		}
	})

}
