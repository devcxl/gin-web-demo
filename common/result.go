package common

type Result struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
