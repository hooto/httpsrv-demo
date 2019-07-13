package main

import (
	"github.com/hooto/httpsrv"

	"github.com/hooto/httpsrv-demo/config"
	"github.com/hooto/httpsrv-demo/websrv/ui"
	"github.com/hooto/httpsrv-demo/websrv/v1"
)

func main() {

	if err := config.Setup(); err != nil {
		panic(err)
	}

	httpsrv.GlobalService.Config.HttpPort = config.Config.HttpPort

	httpsrv.GlobalService.ModuleRegister(config.Config.HttpBasePath+"v1", v1.NewModule())
	httpsrv.GlobalService.ModuleRegister(config.Config.HttpBasePath, ui.NewModule())

	httpsrv.GlobalService.Start()
}
