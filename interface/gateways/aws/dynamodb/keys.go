package dynamodb

import "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

const (
	partitionKey = "pkey"
	sortKey      = "skey"
)

type Key struct {
	Name  string
	Value string
}

type Keys []Key

func (k Keys) SetPartitionKey(value string) Keys {
	return append(k, Key{
		Name:  partitionKey,
		Value: value,
	})
}

func (k Keys) GetPartitionKey() string {
	for _, v := range k {
		if v.Name == partitionKey {
			return v.Value
		}
	}
	return ""
}

func (k Keys) SetSortKey(value string) Keys {
	return append(k, Key{
		Name:  sortKey,
		Value: value,
	})
}

func (k Keys) GetSortKey() string {
	for _, v := range k {
		if v.Name == sortKey {
			return v.Value
		}
	}
	return ""
}

func (k Keys) MarshalAttributeValues() map[string]types.AttributeValue {
	av := make(map[string]types.AttributeValue, len(k))
	for _, v := range k {
		av[v.Name] = &types.AttributeValueMemberS{
			Value: v.Value,
		}
	}
	return av
}
