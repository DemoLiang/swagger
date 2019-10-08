package model
// HttpError example
type HttpError struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"Success"`
}
