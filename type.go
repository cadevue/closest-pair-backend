package main

type SolveCPRequest struct {
	Method    string    `json:"method"`
	Dimension int32     `json:"dimension"`
	Points    []float64 `json:"points"`
}

type SolveCPResponse struct {
	Method            string   `json:"method"`
	Indexes           [2]int32 `json:"indexes"`
	Distance          float64  `json:"distance"`
	NumOfEuclideanOps int32    `json:"numOfEuclideanOps"`
}

type SpecResponse struct {
	CPU    string `json:"cpu"`
	Host   string `json:"host"`
	Memory string `json:"memory"`
	OS     string `json:"os"`
	Disk   string `json:"disk"`
}
