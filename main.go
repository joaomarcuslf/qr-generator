package main

import (
	"os"

	"github.com/joaomarcuslf/qr-generator/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	server := server.NewServer(os.Getenv("PORT"))

	server.Run()
}
