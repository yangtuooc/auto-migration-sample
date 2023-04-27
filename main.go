package main

import (
	"auto-migration-sample/common"
	"auto-migration-sample/routes"
)

func main() {
	startup()
}

func startup() {
	common.Viper()
	common.Gorm()
	routes.Bootstrap()
}
