package rxerrs

import (
	"encoding/json"
)

type ServerError struct {
	msg       string
	status    int
	rootError error
}

func (err ServerError) Error() string {
	return err.msg
}

func (err *ServerError) GetJson() []byte {
	type ServerErrorEvent struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Details string `json:"details"`
	}

	jsonMessage, _ := json.Marshal(ServerErrorEvent{
		Status:  err.status,
		Message: err.msg,
		Details: err.rootError.Error(),
	})

	return jsonMessage
}

func NewServerError(message string, rootError error) ServerError {
	return ServerError{
		msg:       message,
		status:    500,
		rootError: rootError,
	}
}
