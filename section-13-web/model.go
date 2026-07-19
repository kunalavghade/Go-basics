package main

import "time"

type Profile struct {
	UserID int    `json:"user_id"`
	Avatar string `json:"avatar"`
}

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	Profile   Profile   `json:"profile"`
}
