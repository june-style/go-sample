package dynamodb

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/june-style/go-sample/domain/dconfig"
)

func NewClient(cfg *dconfig.Config, awsCfg aws.Config) *Client {
	return &Client{
		db:     dynamodb.NewFromConfig(awsCfg),
		prefix: fmt.Sprintf("%s.%s", cfg.App.Name, cfg.Sys.Env),
	}
}

type Client struct {
	db     *dynamodb.Client
	prefix string
}

func (c *Client) GetDB() *dynamodb.Client {
	return c.db
}

func (c *Client) GetTableFullName(tableName string) string {
	return fmt.Sprintf("%s.%s", c.prefix, tableName)
}
