package model

import "time"

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	PassHash  string    `json:"pass_hash"`
	LastLogin time.Time `json:"last_login"`
}
