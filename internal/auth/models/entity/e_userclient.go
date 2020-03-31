package entity

import "time"

//private key is client id
type UserClient struct {
	ClientId     string    `db:"client_id"`
	ClientSecret string    `db:"client_secret"`
	UserId       string    `db:"user_id"`
	Status       string    `db:"status"`
	CreateTime   time.Time `db:"create_time"`
	UpdateTime   time.Time `db:"update_time"`
}
