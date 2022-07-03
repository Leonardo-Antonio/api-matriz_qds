package response

type response struct {
	MessageType string      `json:"message_type"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

func ResponseErr(
	message string,
	data interface{},
) *response {
	return &response{
		MessageType: ERR,
		Message:     message,
		Data:        data,
	}
}

func ResponseSuccess(
	message string,
	data interface{},
) *response {
	return &response{
		MessageType: SUCCESS,
		Message:     message,
		Data:        data,
	}
}
