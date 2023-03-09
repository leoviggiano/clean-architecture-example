package entity

import (
	"time"
)

type User struct {
	ID   string `db:"id"`
	Name string `db:"name"`

	Exp          int `db:"exp"`
	Level        int `db:"level"`
	NextLevelExp int `db:"next_level_exp"`

	Created time.Time `db:"created_at"`
}
