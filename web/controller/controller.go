package controller

import (
	"github.com/pajlada/pajbot2/pkg/common/config"
	"github.com/pajlada/pajbot2/web/controller/api"
	"github.com/pajlada/pajbot2/web/controller/dashboard"
	"github.com/pajlada/pajbot2/web/controller/home"
	"github.com/pajlada/pajbot2/web/controller/static"
	"github.com/pajlada/pajbot2/web/controller/ws"
)

func LoadRoutes(cfg *config.Config) {
	dashboard.Load()
	home.Load()
	api.Load(cfg)
	static.Load()
	ws.Load()
}
