package response_model

import "github.com/google/uuid"

type Status string

type ErrorResponse struct {
	Errors []ErrorData `json:"errors"`
}

type ErrorData struct {
	ID    uuid.UUID `json:"id"`
	Code  string    `json:"code"`
	Title string    `json:"title"`
}

type BaseResponse[T any] struct {
	Data *T `json:"data"`
}
