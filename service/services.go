package service

import (
	"clean/lib/cache"
	"clean/lib/sql"
)

type All struct {
	User User
}

func GetAll(db sql.Connection, cache cache.Cache) All {
	return All{
		User: NewUser(db, cache),
	}
}
