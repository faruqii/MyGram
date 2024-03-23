package services

type ErrorMessages struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func (e *ErrorMessages) Error() string {
	return e.Message
}

func HandleError(err error, message string, statusCode int) error {
	if err != nil {
		return &ErrorMessages{
			Message:    message,
			StatusCode: statusCode,
		}
	}
	return nil
}
