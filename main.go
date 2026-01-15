package main

import (
	"ToDo/internal/app"
	"ToDo/internal/storage"
)

func main() {

	st := storage.NewStorage()

	todoService := app.NewService(st)

	todoService.Start()
}
