package models

type GeneralResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type RedisReq struct {
	Key  string `json:"key"`
	Data string `json:"data"`
	Exp  int    `json:"exp"`
}
