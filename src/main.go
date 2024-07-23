package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Deck struct {
	Cards       []Card
	Description string
	Size        uint64
}

type Card struct {
	Type    string
	SubType string
}

type TestGame struct{}

type Game interface {
	Start() error
}

type App struct {
	Port string
}

func (a *App) Start() {
	fmt.Print("Deprecating!!!!")
	fmt.Print("Deprecating!!!!")
	fmt.Print("Deprecating!!!!")
	fmt.Print("Deprecating!!!!")

	addr := fmt.Sprintf(":%s", a.Port)
	log.Printf("Starting app on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func env(key, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return val
}

func (t TestGame) Start() {
	fmt.Print("start test game!!")
}

func main() {
	fmt.Print("init main go")

	t := TestGame{}
	t.Start()

	server := App{
		Port: env("PORT", "8080"),
	}

	server.Start()

}
