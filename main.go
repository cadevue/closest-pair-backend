package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  8192,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type SolveCPRequest struct {
	Method string    `json:"method"`
	Points []float64 `json:"points"`
}

type SolveCPResponse struct {
	Method  string   `json:"method"`
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

	log.Printf("Client connected: %s\n", r.RemoteAddr)

	for {
		// Read message from client
		req := SolveCPRequest{}
		err := conn.ReadJSON(&req)
		if err != nil {
			log.Println("Error reading WebSocket request:", err)
			break
		}

		// Solve the closest pair problem
		if req.Method == "dnc" {
			go func() {
				index1, index2 := DnCSolve(req.Points[:])
				sendResponse(req.Method, conn, [2]int32{index1, index2})
			}()
		} else if req.Method == "bruteforce" {
			go func() {
				index1, index2 := BruteforceSolve(req.Points[:])
				sendResponse(req.Method, conn, [2]int32{index1, index2})
			}()
		} else {
			log.Println("Invalid method:", req.Method)
			break
		}
	}

	log.Printf("Client disconnected: %s\n", r.RemoteAddr)
	conn.Close()
}

func sendResponse(method string, conn *websocket.Conn, indexes [2]int32) {
	response := SolveCPResponse{Method: method, Indexes: indexes}
	err := conn.WriteJSON(response)
	if err != nil {
		log.Println("Error writing WebSocket response:", err)
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
