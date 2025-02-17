package main

import (
	"fmt"
	"log"
	"net/http"
)

// "fmt"
// "log"
// "net/http"

func main() {
	// var points []float64 = []float64{
	// 	2, 3, 4,
	// 	1, 2, 3,
	// 	0, 1, 2,
	// 	1, 1, 1,
	// 	4, 4, 4,
	// 	1, 1, 1.1,
	// 	1, 1, 1.4,
	// 	1, 1, 1.7,
	// }
	// result := BruteforceSolve(points, 3)
	// fmt.Printf("Bruteforce: (%d, %d)\n", result.Indexes[0], result.Indexes[1])
	// fmt.Printf("EuclidOps: %d\n", result.NumOfEuclideanOps)

	// fmt.Println()

	// result = DnCSolve(points, 3)
	// fmt.Printf("DnC: (%d, %d)\n", result.Indexes[0], result.Indexes[1])
	// fmt.Printf("EuclidOps: %d\n", result.NumOfEuclideanOps)

	// /*
	http.HandleFunc("/", solveHandler)
	// http.HandleFunc("/spec", specHandler)

	fmt.Println()
	log.Print("Server started on\nws://localhost:8080\n\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
	// */
}
