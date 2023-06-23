package main

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"

	"clean/config"
	"clean/handler"
	"clean/lib/cache"
	"clean/lib/log"
	"clean/lib/sql"
	"clean/service"
)

func main() {
	ctx := context.Background()

	logOptions := make([]log.Option, 0)
	logOptions = append(logOptions, log.WithTextFormatter())
	log := log.NewLogger(logrus.InfoLevel, logOptions...)

	db, err := sql.NewConnection(sql.Settings{
		Conn:            config.DatabaseConnString(),
		MaxIdleCons:     config.MaxIdleConnections(),
		MaxOpenCons:     config.MaxOpenConnections(),
		ConnMaxLifetime: config.MaxLifetimeConnections(),
	}, log)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	cache := cache.StartRedis(ctx, config.RedisAddress(), log)

	services := service.GetAll(db, cache, log)
	handler := handler.NewHandler(services, log)

	r := http.NewServeMux()

	r.HandleFunc("/create-user", handler.CreateUser)
	r.HandleFunc("/delete-user", handler.DeleteUser)
	r.HandleFunc("/get-user", handler.GetUser)
	r.HandleFunc("/get-users", handler.GetUsers)
	r.HandleFunc("/update-user", handler.UpdateUser)

	log.Infof("[Server]: Running on Port %s", config.ServerPort())
	err = http.ListenAndServe(config.ServerPort(), r)
	if err != nil {
		log.Fatal(err)
	}
}
