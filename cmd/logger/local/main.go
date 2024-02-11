package main

import (
	"fmt"

	loggerLocal "github.com/PoorMercymain/go-examples/internal/pkg/logger/logger-local"
)

func main() {
	log, err := loggerLocal.GetLogger()
	if err != nil {
		fmt.Println(err)
		return
	}

	log2, err := loggerLocal.GetLogger()
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Infoln("local logger works!")
	log2.Infoln("local logger 2 works too!")
}
