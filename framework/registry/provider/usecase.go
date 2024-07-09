package provider

import (
	"github.com/google/wire"
	"github.com/june-style/go-sample/application/interactors"
	"github.com/june-style/go-sample/application/usecases"
)

var UsecaseWireSet = wire.NewSet(
	interactors.NewHome,
	interactors.NewSign,
	wire.Struct(new(usecases.UseCase), "*"),
)
