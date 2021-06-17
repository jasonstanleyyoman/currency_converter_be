package utils

type HTTPError struct {
	code int
	message string
	data interface{}
}

func (err * HTTPError) Error() string {
	return err.message
}

func Generate404Error() HTTPError {
	return HTTPError{
		code: 404,
		message: "Error not found",
	}
}