// Package main Notes API server.
//
// @title Notes API
// @version 1.0
// @description Учебный REST API для заметок (CRUD).
// @contact.name Backend Course
// @contact.email example@university.ru
// @BasePath /api/v1
package main

import (
	"log"
	"net/http"

	_ "example.com/notes-api/docs"
	httpapi "example.com/notes-api/internal/http"
	"example.com/notes-api/internal/http/handlers"
	"example.com/notes-api/internal/repo"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	repository := repo.NewNoteRepoMem()
	h := handlers.NewNotesHandler(repository)
	r := httpapi.NewRouter(h)

	r.Get("/docs/*", httpSwagger.WrapHandler)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
