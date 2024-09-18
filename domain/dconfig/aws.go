package dconfig

type Aws struct {
	AccessKeyID     string `env:"AWS_ACCESS_KEY_ID,required"`
	SecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY,required"`
	Region          string `env:"AWS_DEFAULT_REGION,required"`
	Endpoint        string `env:"AWS_ENDPOINT,required"`
	DynamoDBTable   string `env:"AWS_DYNAMO_DB_TABLE,required"`
	DebugMode       bool   `env:"AWS_DEBUG_MODE" envDefault:"false"`
}
