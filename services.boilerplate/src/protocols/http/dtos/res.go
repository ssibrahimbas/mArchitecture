package dtos

import "clean-boilerplate/boilerplate/src/domain/example"

type CreateExampleResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ListExampleResponse struct {
	Offset int               `json:"offset"`
	Limit  int               `json:"limit"`
	Total  int               `json:"total"`
	Items  []example.Example `json:"items"`
}

type GetExampleResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type UpdateExampleResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
