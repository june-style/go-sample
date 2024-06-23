package dynamodb

import "fmt"

type QueryInputData struct {
	Table string
	Key   string
}

func (q QueryInputData) GetTableName() string {
	return q.Table
}

func (q QueryInputData) GetSortKey() Keys {
	return Keys{}.SetSortKey(fmt.Sprintf("%s#", q.Key))
}
