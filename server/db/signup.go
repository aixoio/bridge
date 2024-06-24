package db

import (
	"errors"

	"github.com/aixoio/bridge/server/hasher"
)

func SignUpUser(username, password string) error {
	if len(username) > 32 {
		return errors.New("Username is too long")
	}

	pwd_hash, err := hasher.HashPassword(password)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO users (username, password_hash) VALUES ($1, $2)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(username, pwd_hash)

	return err
}
