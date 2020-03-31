package entity

import "time"

type ClientToken struct {
	TokenId      string    `db:"token_id"`
	ClientId     string    `db:"client_id"`
	RefreshToken string    `db:"refresh_token"`
	Scope        string    `db:"scope"`
	UserId       string    `db:"user_id"`
	Status       string    `db:"status"`
	CreateTime   time.Time `db:"create_time"`
	UpdateTime   time.Time `db:"update_time"`
}
