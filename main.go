package main

import (
	"log"
	"net/http"

	"github.com/AldiRachmatdiantoo/golang-api.git/internal/config"
	"github.com/AldiRachmatdiantoo/golang-api.git/internal/handlers"
	"github.com/AldiRachmatdiantoo/golang-api.git/internal/repositories"
	"github.com/AldiRachmatdiantoo/golang-api.git/internal/router"
	"github.com/AldiRachmatdiantoo/golang-api.git/internal/services"
)

func main() {
	db := config.Connect()
	defer db.Close()
	r := http.NewServeMux()

	repo := repositories.NewRepository(db)
	svc := services.NewService(repo)
	handler := handlers.NewHandler(svc)
	userRouter := router.NewRoute(handler).RegisterRoute()

	r.Handle("/users/", http.StripPrefix("/users", userRouter))

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Fatal(err)
	}

}
