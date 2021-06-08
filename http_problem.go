package httpproblem

import (
	"encoding/json"
	"fmt"
)

// https://tools.ietf.org/html/draft-ietf-appsawg-http-problem-00

// Problem define the struct for http-problem spec
type Problem struct {
	TypeProblem string `json:"type"`
	Title       string `json:"title"`
	Detail      string `json:"detail"`
	Status      int    `json:"status"`
}

// New create problem struct
func New(problemType string, title string, detail string, status int) *Problem {
	var problem Problem

	problem.TypeProblem = problemType
	problem.Title = title
	problem.Detail = detail
	problem.Status = status

	return &problem
}

// Generate build final json
func (problem *Problem) Generate() ([]byte, error) {
	data, _ := json.Marshal(problem)

	return data, nil
}

// GetMediaType return the media type string
func (problem *Problem) GetMediaType() string {
	return "application/api-problem+json"
}

// Error show error from title to adher to error interface
func (problem *Problem) Error() string {
	return fmt.Sprintf("%v | %v", problem.Title, problem.Detail)
}

// GetStatus return the status from the internal struct
func (problem *Problem) GetStatus() int {
	return problem.Status
}

// GetType return probleme type string
func (problem *Problem) GetType() string {
	return problem.TypeProblem
}

// GetDetail return probleme detail
func (problem *Problem) GetDetail() string {
	return problem.Detail
}
