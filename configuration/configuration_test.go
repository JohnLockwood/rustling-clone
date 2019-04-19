package configuration

import (
	"testing"
)

type BadCWD struct {}

func (BadCWD) getCWD() string {
	return "/etc/"
}

func TestBadCWDReturnsAnError(t *testing.T) {
	loader := NewCourseLoader(BadCWD{})
	_, err := loader.GetCourse()
	if err == nil {
		t.Error("Want error, got nil")
	}
}

func TestGoodCWDReturnsNoError(t *testing.T) {
	loader := NewCourseLoader(nil)
	_, err := loader.GetCourse()
	if err != nil {
		t.Errorf("Want nil, got %v", err)
	}
}
