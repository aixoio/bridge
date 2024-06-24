package env

import (
	"os"

	"github.com/joho/godotenv"
)

var ENV Env

func LoadENV() (Env, error) {
	dat, success := loadENVFromSystem()
	if !success {
		dat, err := loadENVFromFile()
		if err != nil {
			return Env{}, nil
		}

		ENV = dat

		return dat, nil
	}

	ENV = dat
	return dat, nil
}

func loadENVFromSystem() (Env, bool) {
	pg_host, exists := os.LookupEnv("PG_HOST")
	if !exists {
		return Env{}, false
	}

	pg_port, exists := os.LookupEnv("PG_PORT")
	if !exists {
		return Env{}, false
	}

	pg_user, exists := os.LookupEnv("PG_USER")
	if !exists {
		return Env{}, false
	}

	pg_pwd, exists := os.LookupEnv("PG_PWD")
	if !exists {
		return Env{}, false
	}

	pg_db, exists := os.LookupEnv("PG_DB")
	if !exists {
		return Env{}, false
	}

	port, exists := os.LookupEnv("PORT")
	if !exists {
		return Env{}, false
	}

	return Env{
		Pg_host: pg_host,
		Pg_port: pg_port,
		Pg_user: pg_user,
		Pg_pwd:  pg_pwd,
		Pg_db:   pg_db,
		Port:    port,
	}, true
}

func loadENVFromFile() (Env, error) {
	err := godotenv.Load()
	if err != nil {
		return Env{}, err
	}

	return Env{
		Pg_host: os.Getenv("PG_HOST"),
		Pg_port: os.Getenv("PG_PORT"),
		Pg_user: os.Getenv("PG_USER"),
		Pg_pwd:  os.Getenv("PG_PWD"),
		Pg_db:   os.Getenv("PG_DB"),
		Port:    os.Getenv("PORT"),
	}, nil
}
