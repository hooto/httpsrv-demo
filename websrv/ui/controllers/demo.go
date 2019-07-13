package controllers

import (
	"github.com/hooto/httpsrv"
)

type Demo struct {
	*httpsrv.Controller
}

func (c Demo) HelloAction() {
	c.RenderString("hello world")
}
