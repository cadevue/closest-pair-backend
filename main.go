package main

import (
	"errors"
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
	Method    string    `json:"method"`
	Dimension int32     `json:"dimension"`
	Points    []float64 `json:"points"`
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

		// Validate the points
		err = isPointsValid(req.Points, int(req.Dimension))
		if err != nil {
			log.Println("Invalid points:", err)
			continue
		}

		log.Printf("Received request: Points : %d, Address : %s, Method : %s\n", len(req.Points)/int(req.Dimension), r.RemoteAddr, req.Method)

		// Solve the closest pair problem
		if req.Method == "dnc" {
			go func() {
				index1, index2 := DnCSolve(req.Points, req.Dimension)
				sendResponse(req.Method, conn, [2]int32{index1, index2})
				log.Printf("Response sent: To : %s, Indexes : (%d, %d), Method : %s\n", r.RemoteAddr, index1, index2, req.Method)
			}()
		} else if req.Method == "bruteforce" {
			go func() {
				index1, index2 := BruteforceSolve(req.Points, req.Dimension)
				sendResponse(req.Method, conn, [2]int32{index1, index2})
				log.Printf("Response sent: To : %s, Indexes : (%d, %d), Method : %s\n", r.RemoteAddr, index1, index2, req.Method)
			}()
		} else {
			log.Println("Invalid method:", req.Method)
			continue
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

func isPointsValid(points []float64, dimension int) error {
	if len(points)%dimension != 0 {
		return errors.New("the number of points is not a multiple of the dimension")
	}

	if len(points) < 2*dimension {
		return errors.New("the number of points is less than 2")
	}

	return nil
}
