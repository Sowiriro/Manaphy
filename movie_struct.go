package main


import "time"

type movie struct {
	Id          int
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
