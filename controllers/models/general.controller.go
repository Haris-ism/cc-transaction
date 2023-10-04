package models

type GeneralResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Token	string		`json:"token,omitempty"`
}