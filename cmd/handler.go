package cmd

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  8192,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return r.Header.Get("Origin") == "https://closest-pair-frontend.pages.dev"
	},
}

func SolveHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket Upgrade Error:", err)
		return
	}
	defer conn.Close()

	log.Printf("Client connected: %s\n", r.RemoteAddr)

	conn.SetReadDeadline(time.Now().Add(90 * time.Second))

	for {
		// Read message from client
		req := SolveCPRequest{}
		err := conn.ReadJSON(&req)
		if err != nil {
			log.Println("Error reading WebSocket request:", err)
			break
		}

		conn.SetReadDeadline(time.Now().Add(90 * time.Second))

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
				result := DnCSolve(req.Points, req.Dimension)
				sendSolveResponse(req.Method, conn, result)
				log.Printf("Response sent: To : %s, Indexes : (%d, %d), Method : %s, Execution Time : %f\n",
					r.RemoteAddr,
					result.Indexes[0], result.Indexes[1],
					req.Method,
					result.ExecutionTime,
				)
			}()
		} else if req.Method == "bruteforce" {
			go func() {
				result := BruteforceSolve(req.Points, req.Dimension)
				sendSolveResponse(req.Method, conn, result)
				log.Printf("Response sent: To : %s, Indexes : (%d, %d), Method : %s, Execution Time : %f\n",
					r.RemoteAddr,
					result.Indexes[0], result.Indexes[1],
					req.Method,
					result.ExecutionTime,
				)
			}()
		} else {
			log.Println("Invalid method:", req.Method)
			continue
		}
	}

	log.Printf("Client disconnected: %s\n", r.RemoteAddr)
	conn.WriteMessage(websocket.CloseMessage, []byte{})
	conn.Close()
}

func sendSolveResponse(method string, conn *websocket.Conn, result SolveResult) {
	response := SolveCPResponse{
		Method:            method,
		Indexes:           result.Indexes,
		Distance:          result.Distance,
		NumOfEuclideanOps: result.NumOfEuclideanOps,
		ExecutionTime:     result.ExecutionTime,
	}
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

	if len(points)/dimension > 100000 {
		return errors.New("the number of points is greater than 100000")
	}

	return nil
}
