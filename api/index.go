package api

import (
	"net/http"

	"github.com/cadevue/closest-pair-backend/cmd"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	cmd.SolveHandler(w, r)
}
