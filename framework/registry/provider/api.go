package provider

import "github.com/google/wire"

var APIWireSet = wire.NewSet(
	ControllerWireSet,
	RepositoryWireSet,
	ServiceWireSet,
	UsecaseWireSet,
)
