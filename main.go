package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

const (
	MAX_NUM_OF_POINTS int32 = 10_000
)

type SolveCPRequest struct {
	Points [MAX_NUM_OF_POINTS]float64 `json:"points"`
}

type SolveCPResponse struct {
	Indexes [2]int32 `json:"indexes"`
}

func main() {
	http.HandleFunc("/", mainHandler)
	fmt.Println()
	log.Print("Server started on\nws://localhost:8080\n\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket Upgrade Error:", err)
		return
	}
	defer conn.Close()

	log.Printf("Client connected: %s\n", r.RemoteAddr)

	// Example: Send a JSON response over WebSocket (not HTTP)
	response := SolveCPResponse{Indexes: [2]int32{0, 1}}
	err = conn.WriteJSON(response)
	if err != nil {
		log.Println("Error writing WebSocket response:", err)
		return
	}
}

/*
Solve closest pair problem using divide and conquer algorithm
returns the index of the closest pair
*/
func DnCSolve(points []float64) (int32, int32) {
	return 0, 1
}

/*
Solve closest pair problem using brute force algorithm
returns the index of the closest pair
*/
func BruteforceSolve(points []float64) (int32, int32) {
	return 0, 1
}
