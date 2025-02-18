package api

import (
	"net/http"

	"github.com/cadevue/closest-pair-backend/internal"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	internal.SolveHandler(w, r)
}
