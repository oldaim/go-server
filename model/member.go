package model

import "time"

type Member struct {
	ID        int64
	Username  string
	Password  string
	Email     string
	Nickname  string
	Status    string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
	LastLogin time.Time
}
