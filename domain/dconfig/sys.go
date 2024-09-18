package dconfig

type Sys struct {
	Env        string `env:"ENVIRONMENT,required"`
	TZ         string `env:"TZ,required"`
	SecretSalt string `env:"SECRET_SALT,required"`
}

const (
	PROD  = "product"
	STG   = "staging"
	DEV   = "development"
	DEV1  = "development1"
	DEV2  = "development2"
	DEV3  = "development3"
	LOCAL = "localhost"
)

func (s Sys) IsProd() bool {
	return s.Env == PROD
}

func (s Sys) IsStg() bool {
	return s.Env == STG
}

func (s Sys) IsDev() bool {
	return s.Env == DEV || s.IsDev1() || s.IsDev2() || s.IsDev3()
}

func (s Sys) IsDev1() bool {
	return s.Env == DEV1
}

func (s Sys) IsDev2() bool {
	return s.Env == DEV2
}

func (s Sys) IsDev3() bool {
	return s.Env == DEV3
}

func (s Sys) IsLocal() bool {
	return s.Env == LOCAL
}
