package controllers

import (
	"github.com/hooto/httpsrv"
	"github.com/hooto/httpsrv-demo/config"
)

type Index struct {
	*httpsrv.Controller
}

func (c Index) IndexAction() {
	c.Data["base_path"] = config.Config.HttpBasePath
}

func (c Index) HelloWorldAction() {
	c.RenderString("hello world")
}
