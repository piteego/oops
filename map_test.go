package oops_test

import (
	"errors"
	"github.com/piteego/oops"
	"os"
	"testing"
)

func TestMap_Handle(t *testing.T) {
	gormErrRecordNotFound := errors.New("gorm record not found")
	redisCacheMissed := errors.New("redis cache missed")
	errMap := oops.Map{
		os.ErrNotExist:        oops.New("file not exists", oops.Tag(NotFound)).(*oops.Error),
		redisCacheMissed:      oops.New("cache key not found", oops.Tag(NotFound)).(*oops.Error),
		gormErrRecordNotFound: oops.New("entity not found", oops.Tag(NotFound)).(*oops.Error),
	}
	for err := range errMap {
		t.Run(err.Error(), func(t *testing.T) {
			got := errMap.Handle(err)
			if !errors.Is(got, err) {
				t.Errorf("expected true, got false")
			}
			if !errors.Is(got, NotFound.Error) {
				t.Errorf("expected true, got false")
			}
		})
	}
	t.Run("duplicated map key", func(t *testing.T) {
		errMap[os.ErrNotExist] = oops.New("file not exists_ duplicated!", oops.Tag(Internal)).(*oops.Error)
		t.Logf("error map with duplicated os.ErrNotExist key: %q", errMap[os.ErrNotExist])
		got := errMap.Handle(os.ErrNotExist)
		if !errors.Is(got, os.ErrNotExist) {
			t.Errorf("expected true, got false")
		}
		if errors.Is(got, NotFound.Error) {
			t.Errorf("expected false, got true")
		}
		if !errors.Is(got, Internal.Error) {
			t.Errorf("expected true, got false")
		}
	})

}
