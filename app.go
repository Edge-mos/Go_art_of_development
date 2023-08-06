package main

import "fmt"

func main() {
	fmt.Println("Error handling")

	err := throwError()
	if err != nil {
		fmt.Println(err)
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

func (ae *AppError) Error() string {
	fmt.Println(ae.Custom)
	return ae.Err.Error()
}

func throwError() error {
	appError := AppError{
		Err:    fmt.Errorf("my error"),
		Custom: "value here",
		Field:  89,
	}
	return &appError
}

//var _ Error = &AppError{}
