package dconfig

type Redis struct {
	Server   string `env:"REDIS_SERVER,required"`
	DBNumber int    `env:"REDIS_DB_NUMBER,required"`
}
