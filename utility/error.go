package utility

import "fmt"

type BadRequestError struct {
	Code    int
	Message string
}

type UnprocessableContentError struct {
	Code    int
	Message string
}

type NotFoundError struct {
	Code    int
	Message string
}

type InternalServerError struct {
	Code    int
	Message string
}

func (err *BadRequestError) Error() string {
	return fmt.Sprintf("BadRequest %v", err.Message)
}

func (err *NotFoundError) Error() string {
	return fmt.Sprintf("NotFound %v", err.Message)
}

func (err *UnprocessableContentError) Error() string {
	return fmt.Sprintf("UnprocessableContent %v", err.Message)
}

func (err *InternalServerError) Error() string {
	return fmt.Sprintf("InternalServerError %v", err.Message)
}
