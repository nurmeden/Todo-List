package cmd

import (
	"log"
	"todo"
	"todo/pkg/handler"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8008", handler.Create_Page); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
