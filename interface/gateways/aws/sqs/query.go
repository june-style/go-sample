package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/june-style/go-sample/domain/dcontext"
	"github.com/june-style/go-sample/domain/derrors"
)

func (c *Client) NewQuery(q QueryInputData) Query {
	return &queryImpl{
		db:        c.GetDB(),
		queueName: c.GetTableFullName(q.GetQueueName()),
	}
}

type Query interface {
	GetQueueURL(ctx context.Context) (string, error)
	GetQueues(ctx context.Context) ([]string, error)
	CreateQueue(ctx context.Context) (string, error)
	DeleteQueue(ctx context.Context) error

	SendMessage(ctx context.Context, message string) (string, error)
	ReceiveMessage(ctx context.Context) (map[string]string, error)
	DeleteMessage(ctx context.Context, receiptHandle string) error
}

type queryImpl struct {
	db        *sqs.Client
	queueName string
}

func (q *queryImpl) GetQueueURL(ctx context.Context) (string, error) {
	params := &sqs.GetQueueUrlInput{
		QueueName: aws.String(q.queueName),
	}
	res, err := q.db.GetQueueUrl(ctx, params)
	if err != nil {
		return "", derrors.Wrap(err)
	}

	return *res.QueueUrl, nil
}

func (q *queryImpl) GetQueues(ctx context.Context) ([]string, error) {
	params := &sqs.ListQueuesInput{}
	res, err := q.db.ListQueues(ctx, params)
	if err != nil {
		return nil, derrors.Wrap(err)
	}

	return res.QueueUrls, nil
}

func (q *queryImpl) CreateQueue(ctx context.Context) (string, error) {
	params := &sqs.CreateQueueInput{
		QueueName: aws.String(q.queueName),
		Attributes: map[string]string{
			"DelaySeconds":           "60",
			"MessageRetentionPeriod": "86400",
		},
	}
	res, err := q.db.CreateQueue(ctx, params)
	if err != nil {
		return "", derrors.Wrap(err)
	}

	return *res.QueueUrl, nil
}

func (q *queryImpl) DeleteQueue(ctx context.Context) error {
	queueURL, err := q.GetQueueURL(ctx)
	if err != nil {
		return derrors.Wrap(err)
	}

	params := &sqs.DeleteQueueInput{
		QueueUrl: aws.String(queueURL),
	}
	if _, err := q.db.DeleteQueue(ctx, params); err != nil {
		return derrors.Wrap(err)
	}

	return nil
}

func (q *queryImpl) ReceiveMessage(ctx context.Context) (map[string]string, error) {
	queueURL, err := q.GetQueueURL(ctx)
	if err != nil {
		return nil, derrors.Wrap(err)
	}

	params := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueURL),
		MaxNumberOfMessages: 1,
		VisibilityTimeout:   10,
	}
	res, err := q.db.ReceiveMessage(ctx, params)
	if err != nil {
		return nil, derrors.Wrap(err)
	}

	messages := make(map[string]string, len(res.Messages))
	for _, set := range res.Messages {
		messages[*set.MessageId] = *set.Body
	}

	return messages, nil
}

func (q *queryImpl) SendMessage(ctx context.Context, message string) (string, error) {
	queueURL, err := q.GetQueueURL(ctx)
	if err != nil {
		return "", derrors.Wrap(err)
	}

	userID := dcontext.GetAuthenticatedUserID(ctx)

	params := &sqs.SendMessageInput{
		MessageBody:  aws.String(message),
		QueueUrl:     aws.String(queueURL),
		DelaySeconds: 10,
		MessageAttributes: map[string]types.MessageAttributeValue{
			"UserID": {
				DataType:    aws.String("String"),
				StringValue: aws.String(userID),
			},
		},
	}
	res, err := q.db.SendMessage(ctx, params)
	if err != nil {
		return "", derrors.Wrap(err)
	}

	return *res.MessageId, nil
}

func (q *queryImpl) DeleteMessage(ctx context.Context, receiptHandle string) error {
	queueURL, err := q.GetQueueURL(ctx)
	if err != nil {
		return derrors.Wrap(err)
	}

	params := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueURL),
		ReceiptHandle: aws.String(receiptHandle),
	}
	if _, err := q.db.DeleteMessage(ctx, params); err != nil {
		return derrors.Wrap(err)
	}

	return nil
}
