package main

import "time"

type Session struct {
	Id         int
	email      string
	user_id    int
	created_at time.Time
}
