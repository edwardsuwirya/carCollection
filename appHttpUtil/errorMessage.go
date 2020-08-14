package appHttpUtil

type ErrorResponseBuilder interface {
	Code() int
	Message() string
}

type ErrorMessage struct {
	ErrorID int    `json:"errorId"`
	Msg     string `json:"message"`
}

func NewErrorMessage(code int, message string) ErrorResponseBuilder {
	return &ErrorMessage{
		ErrorID: code,
		Msg:     message,
	}
}

func (c *ErrorMessage) Code() int {
	return c.ErrorID
}

func (c *ErrorMessage) Message() string {
	return c.Msg
}
