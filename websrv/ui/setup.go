package ui

import (
	"github.com/hooto/httpsrv"

	"github.com/hooto/httpsrv-demo/config"
	"github.com/hooto/httpsrv-demo/websrv/ui/controllers"
)

func NewModule() httpsrv.Module {

	module := httpsrv.NewModule("ui")

	// 设置 View Template 视图模版根路径
	module.TemplatePathSet(config.Prefix + "/websrv/ui/views")

	// 设置前端静态文件访问路径
	module.RouteSet(httpsrv.Route{
		Type:       httpsrv.RouteTypeStatic,
		Path:       "assets",
		StaticPath: config.Prefix + "/webui",
	})

	module.ControllerRegister(new(controllers.Index))
	module.ControllerRegister(new(controllers.Demo))

	return module
}
