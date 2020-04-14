package models

type ErrorValidation struct {
	Message string                 `json:"message"`
	Errors  map[string]interface{} `json:"errors"`
}
