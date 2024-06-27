package validation_test

import (
	"testing"

	"github.com/aixoio/bridge/server/data/validation"
)

func UsernameIsValid_Valid_Test(t *testing.T) {
	if !validation.UsernameIsValid("a-valid-username-1234567890_ASDFGHJKLasdfghjkl") {
		t.Fail()
	}
}

func UsernameIsValid_Invalid_Test(t *testing.T) {
	if !validation.UsernameIsValid("a-invalid-username-~!@#$%^&*()_+`-=[]\\;'\",./<>?") {
		t.Fail()
	}
}
