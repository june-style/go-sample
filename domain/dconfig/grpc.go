package dconfig

type Grpc struct {
	Port        int `env:"GRPC_PORT,required"`
	GatewayPort int `env:"GRPC_GATEWAY_PORT,required"`
}
