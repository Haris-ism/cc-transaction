package models

type GeneralResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Token	string		`json:"token,omitempty"`
}

type ReqHeader struct{
	Authorization	string	`json:"Authorization"`
	TimeStamp		string	`json:"TimeStamp"`
	Signature		string	`json:"Signature"`
}
