package httpproblem

import "encoding/json"

// https://tools.ietf.org/html/draft-ietf-appsawg-http-problem-00

// Problem define the struct for http-problem spec
type Problem struct {
	typeProblem string `json:"type"`
	title       string `json:"title"`
	detail      string `json:"detail"`
	status      int    `json:"status"`
}

// New create problem struct
func New(problemType string, title string, detail string, status int) *Problem {
	var problem Problem

	problem.typeProblem = problemType
	problem.title = title
	problem.detail = detail
	problem.status = status

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
	return problem.title
}
