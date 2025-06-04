package oops_test

import (
	"errors"
	"github.com/piteego/oops"
	"github.com/piteego/oops/example"
	"testing"
)

var (
	gormErrDuplicatedKey  = errors.New("gorm duplicated key")
	gormErrRecordNotFound = errors.New("gorm record not found")
	redisCacheMissed      = errors.New("redis cache missed")
)

func handleRepoErr(entity string) oops.Handler {
	return func(err error) *oops.Error {
		if err == nil {
			return nil
		}
		if errors.Is(err, redisCacheMissed) {
			return oops.New(entity+" not found", oops.Tag(example.NotFound.Error)).(*oops.Error)
		}
		if errors.Is(err, gormErrDuplicatedKey) {
			return oops.New("duplicated "+entity, oops.Tag(example.Duplication.Error)).(*oops.Error)
		}
		if errors.Is(err, gormErrRecordNotFound) {
			return oops.New(entity+" not found", oops.Tag(example.NotFound.Error)).(*oops.Error)
		}
		return oops.New("something went wrong", oops.Tag(example.Internal.Error)).(*oops.Error)
	}
}

func TestHandle(t *testing.T) {
	err := oops.Handle(nil)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	// test cases
	testCases := []struct {
		name     string
		err      error
		expected *oops.Error
	}{
		{"nil error", nil, nil},
		{"redis cache missed", redisCacheMissed, oops.New("entity not found", oops.Tag(example.NotFound.Error)).(*oops.Error)},
		{"gorm duplicated key", gormErrDuplicatedKey, oops.New("duplicated entity", oops.Tag(example.Duplication.Error)).(*oops.Error)},
		{"gorm record not found", gormErrRecordNotFound, oops.New("entity not found", oops.Tag(example.NotFound.Error)).(*oops.Error)},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := handleRepoErr("entity")(tc.err)
			if got == nil && tc.expected == nil {
				return // both are nil, so it's a match
			}
			if got == nil || tc.expected == nil {
				t.Errorf("expected %v, got %v", tc.expected, got)
				return
			}
			if got.Error() != tc.expected.Error() {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
