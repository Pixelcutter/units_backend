package model

import "time"

type UserDetails struct {
	Email    string `json:"email"`
	PassHash string `json:"pass_hash"`
	Username string `json:"username"`
}

type User struct {
	ID        int       `json:"id"`
	Signup    time.Time `json:"signup"`
	LastLogin time.Time `json:"last_login"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
}

var ErrUserExists = "A user with that email or username already exists"
