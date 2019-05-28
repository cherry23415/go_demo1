package main

import (
	"go_gorm_demo/src/server/common"
	"go_gorm_demo/src/server/initialize"
)

func main() {

	initialize.SetupConfig()

	common.SetUpDBGorm()

	initialize.SetErrorDeal()

	initialize.SetContext()

	initialize.SetupServer()
}
