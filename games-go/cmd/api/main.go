package main

import (
	"fmt"
	"games-go-blueprint/internal/database"
	"games-go-blueprint/internal/server"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	fmt.Printf("init server: %s \n", os.Getenv("PORT"))
	server := server.NewServer()

	fmt.Println("init db")
	database.New()
	database.CreateTables()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
	fmt.Print("server should be running")
}
