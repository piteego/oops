package oops_test

import (
	"errors"
	"github.com/piteego/oops"
	"github.com/piteego/oops/example"
	"testing"
)

func TestMap_Handle(t *testing.T) {
	var errMap = oops.Map{
		example.OsErrNotExist:         oops.New("no such file or directory", oops.Tag(example.NotFound.Error)).(*oops.Error),
		example.RedisCacheMissed:      oops.New("cache key not found", oops.Tag(example.NotFound.Error)).(*oops.Error),
		example.GormErrRecordNotFound: oops.New("entity not found", oops.Tag(example.NotFound.Error)).(*oops.Error),
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
		errMap[example.OsErrNotExist] = oops.New(
			"no such file or directory__duplicated!",
			oops.Tag(example.Internal.Error),
		).(*oops.Error)
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
