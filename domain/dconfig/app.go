package dconfig

type App struct {
	Name                  string `env:"APPLICATION_NAME,required"`
	Key                   string `env:"APPLICATION_KEY,required"`
	SessionExpirationTime int    `env:"APPLICATION_SESSION_EXPIRATION_TIME,required"`
	HMACSecret            string `env:"APPLICATION_HMAC_SECRET,required"`
}
