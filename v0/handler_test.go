package v0_test

import (
	"github.com/piteego/oops/example"
	"github.com/piteego/oops/v0"
	"testing"
)

func TestHandle(t *testing.T) {
	err := v0.Handle(nil)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	// test cases
	testCases := []struct {
		name     string
		err      error
		expected *v0.Error
	}{
		{"nil error", nil, nil},
		{"redis cache missed", example.RedisCacheMissed, v0.New("entity not found", v0.Tag(example.NotFound.Error)).(*v0.Error)},
		{"gorm duplicated key", example.GormErrDuplicatedKey, v0.New("duplicated entity", v0.Tag(example.Duplication.Error)).(*v0.Error)},
		{"gorm record not found", example.GormErrRecordNotFound, v0.New("entity not found", v0.Tag(example.NotFound.Error)).(*v0.Error)},
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
