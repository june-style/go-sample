package redis

import (
	"fmt"

	"github.com/june-style/go-sample/domain/dconfig"
	"github.com/june-style/go-sample/domain/derrors"
	"github.com/redis/go-redis/v9"
)

func NewClient(cfg *dconfig.Config) (*Client, error) {
	dbMap := make(dbMap, cfg.Redis.DBNumber)
	for i := 0; i < cfg.Redis.DBNumber; i++ {
		dbMap[DBIndex(i)] = redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs: []string{cfg.Redis.Server},
			DB:    i,
		})
	}
	return &Client{
		dbMap:  dbMap,
		prefix: fmt.Sprintf("%s.%s", cfg.App.Name, cfg.Sys.Env),
	}, nil
}

type Client struct {
	dbMap  dbMap
	prefix string
}

func (c *Client) GetDB(idx DBIndex) (redis.UniversalClient, error) {
	if db, ok := c.dbMap[idx]; ok {
		return db, nil
	}
	return nil, derrors.Wrapf(ErrNotFoundRedisDBByIndex, "index is %d", idx)
}

func (c *Client) GetKeyFullName(keyName string) string {
	return fmt.Sprintf("%s.%s", c.prefix, keyName)
}

type dbMap map[DBIndex]redis.UniversalClient

type DBIndex int

const (
	DB_00 DBIndex = iota
	DB_01
	DB_02
	DB_03
	DB_04
	DB_05
	DB_06
	DB_07
	DB_08
	DB_09
	DB_10
	DB_11
	DB_12
	DB_13
	DB_14
	DB_15
)

var ErrNotFoundRedisDBByIndex = derrors.NewInternal("not found redis db by index")
