package main

type SolveCPRequest struct {
	Method    string    `json:"method"`
	Dimension int32     `json:"dimension"`
	Points    []float64 `json:"points"`
}

type SolveCPResponse struct {
	Method  string   `json:"method"`
	Indexes [2]int32 `json:"indexes"`
}
