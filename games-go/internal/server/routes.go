package server

import (
	"encoding/json"
	"games-go-blueprint/internal/database"
	"games-go-blueprint/internal/models"
	"log"
	"net/http"
	"text/template"

	"fmt"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"nhooyr.io/websocket"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fileServer := http.FileServer(http.Dir("./assets"))
	r.Handle("/assets/*", http.StripPrefix("/assets", fileServer))

	r.Get("/", s.AppHandler)

	r.Get("/health", s.healthHandler)
	r.Get("/test", s.testHandler)
	// r.Get("/lobby", s.)

	r.Get("/websocket", s.websocketHandler)

	return r
}

func (s *Server) AppHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
	// w.W
	// fs := http.FileServer(http.Dir("assets"))
	// // fs2 := http.FileServer(http.Dir("static/assets"))
	// // r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// http.StripPrefix("/", fs).ServeHTTP(w, r)

	// http.StripPrefix("/assets", fs2).ServeHTTP(w, r)
	// })
}

// fileServer := http.FileServer(http.Dir("./assets"))
// r.Handle("/assets/*", http.StripPrefix("/assets", fileServer))
func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) testHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetConnection()

	_, err := models.AddUser(db, "devKev", "test@test.de")
	if err != nil {
		log.Fatal(err)
	}

}

func (s *Server) gameHandler(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}

func (s *Server) websocketHandler(w http.ResponseWriter, r *http.Request) {
	socket, err := websocket.Accept(w, r, nil)

	if err != nil {
		log.Printf("could not open websocket: %v", err)
		_, _ = w.Write([]byte("could not open websocket"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer socket.Close(websocket.StatusGoingAway, "server closing websocket")

	ctx := r.Context()
	socketCtx := socket.CloseRead(ctx)

	for {
		payload := fmt.Sprintf("server timestamp: %d", time.Now().UnixNano())
		err := socket.Write(socketCtx, websocket.MessageText, []byte(payload))
		if err != nil {
			break
		}
		time.Sleep(time.Second * 2)
	}
}
