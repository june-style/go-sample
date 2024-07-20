package provider

import (
	"github.com/google/wire"
	"github.com/june-style/go-sample/interface/controllers"
)

var ControllerWireSet = wire.NewSet(
	controllers.NewHome,
	controllers.NewSign,
	wire.Struct(new(controllers.Controller), "*"),
)
