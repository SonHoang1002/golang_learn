package models

import "time"

type UserModel struct {
	Username  string
	Age       int
	CreatedAt time.Time
}
