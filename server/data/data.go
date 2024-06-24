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

type Transaction struct {
	Id                int
	Recipient_user_id int
	Sender_user_id    int
	Is_mined          bool
	Sent_at           time.Time
	Tx_description    string
}
