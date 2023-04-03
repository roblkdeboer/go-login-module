package errors

type BadRequestError struct {
	Message string
}

func (e BadRequestError) Error() string {
	return e.Message
}

type DatabaseError struct {
	Message string
}

func (e DatabaseError) Error() string {
	return e.Message
}