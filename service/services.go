package service

import (
	"clean/lib/cache"
	"clean/lib/log"
	"clean/lib/sql"
)

type All struct {
	User User
}

func GetAll(db sql.Connection, cache cache.Cache, log log.Logger) All {
	return All{
		User: NewUser(db, cache, log),
	}
}
