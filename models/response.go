package models

// Response -> model
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
