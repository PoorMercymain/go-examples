package main

import (
	"errors"
	"fmt"
	"log"
)

type IMyErr interface {
	error
	GetVal() int
}

type MyErr struct{}

func (m MyErr) Error() string {
	return "an error"
}

func (m MyErr) GetVal() int {
	return 1
}

func main() {
	err := errors.New("some error")
	log.Println(err)

	wrapErr := fmt.Errorf("some error happened: %w", err)
	log.Println(wrapErr)

	err = errors.New("Some error")
	wrapErr = fmt.Errorf("some error happened: %w", err)
	log.Println(wrapErr)

	err = errors.New("some error")
	wrapErr = fmt.Errorf("some error happened: %w", err)

	if unwrappedErr := errors.Unwrap(wrapErr); unwrappedErr != nil {
		log.Println("successfully unwrapped wrapErr:", unwrappedErr)
	}

	notWrappedErr := fmt.Errorf("some error happened but is not wrapped: %v", err)
	if unwrappedErr := errors.Unwrap(notWrappedErr); unwrappedErr != nil {
		log.Println("successfully unwrapped notWrappedErr:", unwrappedErr)
	}

	my := MyErr{}

	err = fmt.Errorf("error: %w", my)

	var myNew IMyErr

	if errors.As(err, &myNew) {
		log.Println("errors.As succeeded:", myNew)
	}

	if errors.As(errors.New("random err"), &myNew) {
		log.Println("errors.As succeeded for wrong case:", myNew)
	}
}
