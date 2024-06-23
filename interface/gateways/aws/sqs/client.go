package sqs

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/june-style/go-sample/domain/dconfig"
)

func NewClient(cfg *dconfig.Config, awsCfg aws.Config) *Client {
	return &Client{
		db:     sqs.NewFromConfig(awsCfg),
		prefix: fmt.Sprintf("%s.%s", cfg.App.Name, cfg.Sys.Env),
	}
}

type Client struct {
	db     *sqs.Client
	prefix string
}

func (c *Client) GetDB() *sqs.Client {
	return c.db
}

func (c *Client) GetTableFullName(tableName string) string {
	return fmt.Sprintf("%s.%s", c.prefix, tableName)
}
