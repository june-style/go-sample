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

type UserProfile struct {
	UserID    string    `dynamodbav:"pkey"`
	SortKey   string    `dynamodbav:"skey"`
	Name      string    `dynamodbav:"name"`
	CreatedAt time.Time `dynamodbav:"created_at"`
}

type UserProfileSet []*UserProfile

func NewUserProfileRepository(cfg *dconfig.Config, client *aws.Client) (entities.UserProfileRepository, error) {
	input := dynamodb.QueryInputData{
		Table: cfg.Aws.DynamoDBTable,
		Key:   "profile",
	}
	return &userProfileImpl{
		query: client.DynamoDB.NewQuery(input),
		keys:  input.GetSortKey(),
	}, nil
}

type userProfileImpl struct {
	query dynamodb.Query
	keys  dynamodb.Keys
}

func (u *userProfileImpl) Find(ctx context.Context, userID string) (*entities.UserProfile, error) {
	keys := u.keys.SetPartitionKey(userID)
	profile := &UserProfile{}
	if err := u.query.Get(ctx, keys, profile); err != nil {
		return nil, derrors.Wrap(err)
	}
	return entities.NewUserProfile(
		profile.UserID,
		profile.Name,
		profile.CreatedAt,
	), nil
}

func (u *userProfileImpl) Create(ctx context.Context, userProfile *entities.UserProfile) error {
	if err := u.query.Put(ctx, &UserProfile{
		UserID:    userProfile.UserID(),
		SortKey:   u.keys.GetSortKey(),
		Name:      userProfile.Name(),
		CreatedAt: userProfile.CreatedAt(),
	}); err != nil {
		return derrors.Wrap(err)
	}
	return nil
}

func (u *userProfileImpl) Delete(ctx context.Context, userProfile *entities.UserProfile) error {
	keys := u.keys.SetPartitionKey(userProfile.UserID())
	if err := u.query.Del(ctx, keys); err != nil {
		return derrors.Wrap(err)
	}
	return nil
}
