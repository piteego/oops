package oops_test

import (
	"errors"
	"github.com/piteego/oops"
	"github.com/piteego/oops/example"
	"os"
	"testing"
)

func TestMap_Handle(t *testing.T) {
	var errMap = oops.Map{
		os.ErrNotExist:        oops.New("file not exists", oops.Tag(example.NotFound.Error)).(*oops.Error),
		redisCacheMissed:      oops.New("cache key not found", oops.Tag(example.NotFound.Error)).(*oops.Error),
		gormErrRecordNotFound: oops.New("entity not found", oops.Tag(example.NotFound.Error)).(*oops.Error),
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
		errMap[os.ErrNotExist] = oops.New("file not exists_ duplicated!", oops.Tag(example.Internal.Error)).(*oops.Error)
		t.Logf("error map with duplicated os.ErrNotExist key: %q", errMap[os.ErrNotExist])
		got := errMap.Handle(os.ErrNotExist)
		if !errors.Is(got, os.ErrNotExist) {
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
