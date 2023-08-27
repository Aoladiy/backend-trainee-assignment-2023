package main

import (
	todo "github.com/Aoladiy/backend-trainee-assignment-2023"
	"github.com/Aoladiy/backend-trainee-assignment-2023/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	server := new(todo.Server)
	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Http server error %v", err)
	}

}
