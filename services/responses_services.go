package services

const (
	Error   = "error"
	Message = "message"
)

type response struct {
	MessageType string      `json:"message_type"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}
