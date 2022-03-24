package rxerrs

import (
	"encoding/json"
)

type UserError struct {
	msg    string
	status int
}

func (err UserError) Error() string {
	return err.msg
}

func (err *UserError) GetJson() []byte {
	type UserErrorEvent struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	jsonMessage, _ := json.Marshal(UserErrorEvent{
		Status:  err.status,
		Message: err.msg,
	})

	return jsonMessage
}

func NewUserError(message string, status int) UserError {
	return UserError{
		msg:    message,
		status: status,
	}
}
