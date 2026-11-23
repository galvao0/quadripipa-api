package main

import (
	"github.com/galvao0/quadripipa-api/config"
	"github.com/galvao0/quadripipa-api/router"
)

func main() {

	logger := config.GetLogger("main")

	if err := config.Init(); err != true {
		logger.Error(err)
		return
	}

	router.Init()

}
