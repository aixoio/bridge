package validation

import "regexp"

func UsernameIsValid(username string) bool {
	return regexp.MustCompile("^[A-Za-z0-9_-]+$").MatchString(username)
}

func TransactionValueIsValid(value int64) bool {
	return value > 0
}
