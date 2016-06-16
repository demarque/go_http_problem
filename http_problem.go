package httpproblem

import "encoding/json"

// https://tools.ietf.org/html/draft-ietf-appsawg-http-problem-00

// Problem define the struct for http-problem spec
type Problem struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Status int    `json:"status"`
}

// New create problem struct
func New(problemType string, title string, detail string, status int) Problem {
	var problem Problem

	problem.Type = problemType
	problem.Title = title
	problem.Detail = detail
	problem.Status = status

	return problem
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
