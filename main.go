package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cadevue/closest-pair-backend/internal"

	"github.com/didip/tollbooth/v8"
	"github.com/didip/tollbooth/v8/limiter"
)

func main() {
	// Rate limiting
	lmt := tollbooth.NewLimiter(100, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})
	lmt.SetIPLookup(limiter.IPLookup{
		Name:           "X-Forwarded-For",
		IndexFromRight: 0,
	})
	lmt.SetMessage("You have reached maximum request limit per hour.")

	http.Handle("/", tollbooth.LimitHandler(lmt, http.HandlerFunc(internal.SolveHandler)))

	fmt.Println()
	log.Print("Server started on\nws://localhost:8080\n\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
