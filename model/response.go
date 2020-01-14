package model

import "net/http"

const SuccessCode = 0

const SuccessMsg = "ok"

type JsonResponse struct {
	ErrorCode   int         `json:"errcode"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

func NewSuccessfulJsonResponse(data interface{}) *JsonResponse {
	resp := JsonResponse{
		ErrorCode:   SuccessCode,
		Description: SuccessMsg,
		Data:        data,
	}
	return &resp
}

func NewInternalServerErrorJsonResponse(err error) *JsonResponse {
	resp := JsonResponse{
		ErrorCode:   http.StatusInternalServerError,
		Description: err.Error(),
	}
	return &resp
}

type BlockchainResponse struct {
	TxID     string `json:"tx_id"`
	Response string `json:"response"`
}
