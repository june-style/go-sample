package provider

import (
	"github.com/google/wire"
	"github.com/june-style/go-sample/domain/services"
)

var ServiceWireSet = wire.NewSet(
	services.NewAuthorizer,
	services.NewJWTer,
	services.NewTimer,
	wire.Struct(new(services.Service), "*"),
)
