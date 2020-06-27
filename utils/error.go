package utils

type Error struct {
	message string "message:"
	Data    interface{}
}

func GenerateError(msg string, data interface{}) Error {
	return Error{
		message: msg,
		Data:    data,
	}
}
