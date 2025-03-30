package models

// BaseResponse 基础响应结构
type BaseResponse struct {
	Code      int         `json:"code"`
	RequestID string      `json:"request_id"`
	Message   string      `json:"message"`
	Time      int64       `json:"time"`
	Usage     int         `json:"usage"`
	Data      interface{} `json:"data"`
}
