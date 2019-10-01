//package aws_sqs
package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

type mockedReceiveMsgs struct {
	sqsiface.SQSAPI
	Resp sqs.ReceiveMessageOutput
}

//
//func (m *mockedReceiveMsgs) ReceiveMessage(in *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
//	return &m.Resp, nil
//}
//
//func (m *mockedReceiveMsgs) Receive413Error(in *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
//	return &m.Resp, fmt.Errorf("413 error")
//}

func (q *Queue) GetMessages(waitTimeout int64) ([]MockMessage, error) {
	params := sqs.ReceiveMessageInput{
		QueueUrl: aws.String(q.URL),
	}
	if waitTimeout > 0 {
		params.WaitTimeSeconds = aws.Int64(waitTimeout)
	}
	resp, err := q.Client.ReceiveMessage(&params)

	if err != nil {
		return nil, err
	}

	msgs := make([]MockMessage, len(resp.Messages))
	for i, msg := range resp.Messages {
		parsedMsg := MockMessage{}
		if err := json.Unmarshal([]byte(aws.StringValue(msg.Body)), &parsedMsg); err != nil {
			return nil, fmt.Errorf("failed to unmarshal message, %v", err)
		}

		msgs[i] = parsedMsg
	}
	return msgs, nil
}

// Message is a concrete representation of the SQS message
type MockMessage struct {
	From string `json:"from"`
	To   string `json:"to"`
	Msg  string `json:"msg"`
}

type Queue struct {
	Client             sqsiface.SQSAPI
	URL                string
	MaximumMessageSize int
}
