package main

import (
	"fmt"

	loggerInit "github.com/PoorMercymain/go-examples/internal/pkg/logger/logger-init"
	loggerOnce "github.com/PoorMercymain/go-examples/internal/pkg/logger/logger-once"
)

func main() {
	loggerOnce.Logger().Infoln("logger works!")

	err := loggerInit.InitLogger()
	if err != nil {
		fmt.Println(err)
		return
	}

	log, err := loggerInit.GetLogger()
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Infoln("init logger works too!")
}
