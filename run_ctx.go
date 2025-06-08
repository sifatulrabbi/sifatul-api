package sifatulapi

type RunCtxKey string

const (
	GOENV   RunCtxKey = "GOENV"
	PORT    RunCtxKey = "PORT"
	DB_CONN RunCtxKey = "DB_CONN"
)
