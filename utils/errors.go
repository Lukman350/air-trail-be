package utils

type NotFoundError struct {
	Message string
}

type ValidationError struct {
	Message string
}

func (err *ValidationError) Error() string {
	return err.Message
}

func (err *NotFoundError) Error() string {
	return err.Message
}

type Cat021Error struct {
	Message string
}

func (err *Cat021Error) Error() string {
	return err.Message
}
