package v1

import (
	"github.com/hooto/httpsrv"
)

func NewModule() httpsrv.Module {

	module := httpsrv.NewModule("api")

	module.ControllerRegister(new(User))

	return module
}
