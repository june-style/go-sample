package redis

import "time"

type QueryInputData struct {
	KeyName    string
	DBIndex    DBIndex
	ExpireTime time.Duration
}

func (q QueryInputData) GetKeyName() string {
	return q.KeyName
}

func (q QueryInputData) GetDBIndex() DBIndex {
	return q.DBIndex
}

func (q QueryInputData) GetExpireTime() time.Duration {
	return q.ExpireTime
}
