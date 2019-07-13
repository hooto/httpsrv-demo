package v1

import (
	"github.com/hooto/httpsrv"
)

type User struct {
	*httpsrv.Controller
}

func (c User) InfoAction() {

	jsonStruct := struct {
		Name string `json:"name"`
	}{
		Name: "robot",
	}

	c.RenderJson(&jsonStruct)
}
