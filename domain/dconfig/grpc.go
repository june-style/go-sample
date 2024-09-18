package dconfig

type Grpc struct {
	Port int `env:"GRPC_PORT,required"`
}
