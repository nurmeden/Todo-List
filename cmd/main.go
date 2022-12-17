package cmd

import (
	"log"
	"todo"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8008"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
