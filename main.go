package main

import "fmt"

// "fmt"
// "log"
// "net/http"

func main() {
	var points []float64 = []float64{2, 3, 4, 1, 2, 3, 0, 1, 2, 1, 1, 1, 4, 4, 4}
	b1, b2 := BruteforceSolve(points, 3)
	fmt.Printf("Bruteforce: (%d, %d)\n", b1, b2)

	d1, d2 := DnCSolve(points, 3)
	fmt.Printf("DnC: (%d, %d)\n", d1, d2)

	/*
		http.HandleFunc("/", solveHandler)
		http.HandleFunc("/spec", specHandler)

		fmt.Println()
		log.Print("Server started on\nws://localhost:8080\n\n")
		log.Fatal(http.ListenAndServe(":8080", nil))
	*/
}
