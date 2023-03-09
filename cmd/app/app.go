package main

import (
	"context"
	"log"
	"net/http"

	"clean/handler"
	"clean/lib/cache"
	"clean/lib/sql"
	"clean/service"
)

func main() {
	ctx := context.Background()

	cache := cache.StartRedis(ctx, "conexão do redis")
	db, err := sql.NewConnection("aquela string grande de conectar no banco")
	if err != nil {
		log.Fatal(err)
	}

	services := service.GetAll(db, cache)
	handler := handler.NewHandler(services)

	r := http.NewServeMux()

	r.HandleFunc("/create-user", handler.CreateUser)
	r.HandleFunc("/delete-user", handler.DeleteUser)
	r.HandleFunc("/get-user", handler.GetUser)
	r.HandleFunc("/update-user", handler.UpdateUser)

	err = http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}
}
