package sqs

type QueryInputData struct {
	QueueName string
}

func (q QueryInputData) GetQueueName() string {
	return q.QueueName
}
