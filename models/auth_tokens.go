package models

type RevokedToken struct {
	Token string `db:"token"`
}
