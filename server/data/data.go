package data

import "time"

type User struct {
	Id             int
	Username       string
	Password_hash  string
	Balance        int64
	Public_Account bool
	Created_At     time.Time
}
