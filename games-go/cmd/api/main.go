package main

import (
	"fmt"
	"games-go-blueprint/internal/database"
	"games-go-blueprint/internal/models"
	"games-go-blueprint/internal/server"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	fmt.Printf("now: %s \n", time.DateTime)
	fmt.Printf("init server: %s \n", os.Getenv("PORT"))
	server := server.NewServer()

	service := database.New()
	if service == nil {
		log.Fatal("couldnt get db")
	}

	models.CreateUserTable(service.GetConnection())

	fmt.Print("tables and co should be initialized")

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
	fmt.Print("server should be running")

	// database.

}
