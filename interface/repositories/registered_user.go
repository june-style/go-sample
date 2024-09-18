package repositories

import (
	"context"
	"time"

	"github.com/june-style/go-sample/domain/dconfig"
	"github.com/june-style/go-sample/domain/derrors"
	"github.com/june-style/go-sample/domain/entities"
	"github.com/june-style/go-sample/interface/gateways/aws"
	"github.com/june-style/go-sample/interface/gateways/aws/dynamodb"
)

type RegisteredUser struct {
	AccessKey string    `dynamodbav:"pkey"`
	SortKey   string    `dynamodbav:"skey"`
	UserID    string    `dynamodbav:"user_id"`
	CreatedAt time.Time `dynamodbav:"created_at"`
}

type RegisteredUserSet []*RegisteredUser

func NewRegisteredUserRepository(cfg *dconfig.Config, client *aws.Client) (entities.RegisteredUserRepository, error) {
	input := dynamodb.QueryInputData{
		Table: cfg.Aws.DynamoDBTable,
		Key:   "registry",
	}
	return &registeredUserImpl{
		query: client.DynamoDB.NewQuery(input),
		keys:  input.GetSortKey(),
	}, nil
}

type registeredUserImpl struct {
	query dynamodb.Query
	keys  dynamodb.Keys
}

func (d *registeredUserImpl) Find(ctx context.Context, userID string) (*entities.RegisteredUser, error) {
	keys := d.keys.SetPartitionKey(userID)
	register := &RegisteredUser{}
	if err := d.query.Get(ctx, keys, register); err != nil {
		return nil, derrors.Wrap(err)
	}
	return entities.NewRegisteredUser(
		register.AccessKey,
		register.UserID,
		register.CreatedAt,
	), nil
}

func (d *registeredUserImpl) Create(ctx context.Context, registeredUser *entities.RegisteredUser) error {
	if err := d.query.Put(ctx, &RegisteredUser{
		AccessKey: registeredUser.AccessKey(),
		SortKey:   d.keys.GetSortKey(),
		UserID:    registeredUser.UserID(),
		CreatedAt: registeredUser.CreatedAt(),
	}); err != nil {
		return derrors.Wrap(err)
	}
	return nil
}

func (d *registeredUserImpl) Delete(ctx context.Context, registeredUser *entities.RegisteredUser) error {
	keys := d.keys.SetPartitionKey(registeredUser.AccessKey())
	if err := d.query.Del(ctx, keys); err != nil {
		return derrors.Wrap(err)
	}
	return nil
}
