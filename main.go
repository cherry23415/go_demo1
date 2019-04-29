package main

import (
	"go_demo1/src/server/common"
	"go_demo1/src/server/initialize"
)

func main() {

	initialize.SetupConfig()

	common.SetUpDBGorm()

	initialize.SetErrorDeal()

	initialize.SetContext()

	initialize.SetupServer()
}
