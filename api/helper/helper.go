package helper

import (
	"encoding/json"
	"net/http"
)

type JsonResponse struct {
	HttpStatus int
	RW         http.ResponseWriter
	Encoder    *json.Encoder
}

func NewJsonResponse(w http.ResponseWriter) *JsonResponse {
	jr := JsonResponse{RW: w, Encoder: json.NewEncoder(w)}
	jr.RW.Header().Set("Content-Type", "application/json")
	return &jr
}

func (jr *JsonResponse) SetStatus(st int) *JsonResponse {
	jr.HttpStatus = st
	jr.RW.WriteHeader(st)
	return jr
}

func (jr *JsonResponse) Encode(v interface{}) *JsonResponse {
	jr.Encoder.Encode(v)
	return jr
}
