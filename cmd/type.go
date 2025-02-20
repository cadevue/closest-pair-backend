package cmd

/* Request-Response */
type SolveCPRequest struct {
	Method    string    `json:"method"`
	Dimension int32     `json:"dimension"`
	Points    []float64 `json:"points"`
}

type SolveCPResponse struct {
	Method            string   `json:"method"`
	Indexes           [2]int32 `json:"indexes"`
	Distance          float64  `json:"distance"`
	NumOfEuclideanOps int64    `json:"numOfEuclideanOps"`
	ExecutionTime     float64  `json:"executionTime"`
}

type SpecResponse struct {
	CPU    string `json:"cpu"`
	Host   string `json:"host"`
	Memory string `json:"memory"`
	OS     string `json:"os"`
	Disk   string `json:"disk"`
}

/* Internal Data */
type SolveData struct {
	Points            []float64
	Dimension         int32
	NumOfEuclideanOps int64
}

type SolveResult struct {
	Indexes           [2]int32
	Distance          float64
	NumOfEuclideanOps int64
	ExecutionTime     float64
}
