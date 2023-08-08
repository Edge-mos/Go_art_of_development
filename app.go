package main

import "fmt"

func main() {
	fmt.Println("Error handling")

	err := throwError()
	if err != nil {
		fmt.Println(err)
		err.Unwrap()
	}

}

//type Error interface {
//	Error() string
//}

type AppError struct {
	Err    error
	Custom string
	Field  int
}

func (ae *AppError) Error() string { // реализуем тут интерфейс error
	fmt.Println(ae.Custom)
	return ae.Err.Error()
}

func (ae *AppError) Unwrap() error {
	return ae.Err
}

func throwError() *AppError {
	appError := AppError{
		Err:    fmt.Errorf("my error"),
		Custom: "value here",
		Field:  89,
	}
	return &appError
}

var _ error = &AppError{}
