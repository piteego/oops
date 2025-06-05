package oops_test

import (
	"github.com/piteego/oops"
	"github.com/piteego/oops/example"
	"testing"
)

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
		{"redis cache missed", example.RedisCacheMissed, oops.New("entity not found", oops.Tag(example.NotFound.Error)).(*oops.Error)},
		{"gorm duplicated key", example.GormErrDuplicatedKey, oops.New("duplicated entity", oops.Tag(example.Duplication.Error)).(*oops.Error)},
		{"gorm record not found", example.GormErrRecordNotFound, oops.New("entity not found", oops.Tag(example.NotFound.Error)).(*oops.Error)},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := example.HandleRepoErr("entity")(tc.err)
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
