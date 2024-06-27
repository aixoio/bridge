package validation_test

import (
	"testing"

	"github.com/aixoio/bridge/server/data/validation"
)

func TestUsernameIsValidValid(t *testing.T) {
	if !validation.UsernameIsValid("a-valid-username-1234567890_ASDFGHJKLasdfghjkl") {
		t.Fail()
	}
}

func TestUsernameIsValidInvalid(t *testing.T) {
	if validation.UsernameIsValid("a-invalid-username-~!@#$%^&*()_+`-=[]\\;'\",./<>?") {
		t.Fail()
	}
}

func TestTransactionValueIsValidValid(t *testing.T) {
	if !validation.TransactionValueIsValid(8) {
		t.Fail()
	}
}

func TestTransactionValueIsValidInvalid(t *testing.T) {
	if validation.TransactionValueIsValid(-4) {
		t.Fail()
	}

	if validation.TransactionValueIsValid(0) {
		t.Fail()
	}
}
