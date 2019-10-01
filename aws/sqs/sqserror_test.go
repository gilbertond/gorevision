//package aws_sqs
package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"log"
	"testing"

	//"github.com/aws/aws-sdk-go/aws"
	//"github.com/aws/aws-sdk-go/aws/session"
	//"github.com/aws/aws-sdk-go/service/sqs"
	"os"
)

func TestSQS(t *testing.T) {

	//tooBig := string(make([]byte, 1048577))
	//sender :=
	RegisterFailHandler(Fail)
	RunSpecs(t, "SQS Test Suite")
}

var _ = Describe("SQS Tests", func() {

	//var myCall *gomock.Call

	//var mockReceiveMessages *mockedReceiveMsgs

	BeforeEach(func() {
		sess := session.Must(session.NewSession(&aws.Config{Region:aws.String("us-east-2")}))
		q := Queue{
			Client: sqs.New(sess),
			URL:    "http://localhost:9292/drain",
			MaximumMessageSize:100,
		}

		msgs, err := q.GetMessages(20)
		Expect(err).NotTo(HaveOccurred())

		fmt.Println("Messages:")
		for _, msg := range msgs {
			fmt.Printf("%s>%s: %s\n", msg.From, msg.To, msg.Msg)
		}
	})

	AfterEach(func() {

	})

	Context("Success 😎", func() {
		//It("send a message return with no error", func() {
		//	mockReceiveMessages, error := mockReceiveMessages.ReceiveMessage(&inputMessage)
		//	Expect(error).To(HaveOccurred())
		//	Expect(mockReceiveMessages).To(BeNil())
		//})

		It("get Message", func() {
			cases := []struct {
				Resp     sqs.ReceiveMessageOutput
				Expected []MockMessage
			}{
				{
					Resp: sqs.ReceiveMessageOutput{
						Messages: []*sqs.Message{
							{Body: aws.String(`{"from":"user_1","to":"room_1","msg":"Hello!"}`)},
							{Body: aws.String(`{"from":"user_2","to":"room_1","msg":"Hi user_1 :)"}`)},
						},
					},
					Expected: []MockMessage{
						{From: "user_1", To: "room_1", Msg: "Hello!"},
						{From: "user_2", To: "room_1", Msg: "Hi user_1 :)"},
					},
				},
				{ // Case 2, not messages returned
					Resp:     sqs.ReceiveMessageOutput{},
					Expected: []MockMessage{},
				},
			}

			for i, c := range cases {
				q := Queue{
					Client: mockedReceiveMsgs{Resp: c.Resp},
					URL:    fmt.Sprintf("mockURL_%d", i),
				}
				msgs, err := q.GetMessages(20)
				if err != nil {
					log.Fatalf("%d, unexpected error", err)
				}
				if a, e := len(msgs), len(c.Expected); a != e {
					log.Fatalf("%d, expected %d messages, got %d", i, e, a)
				}
				//for j, msg := range msgs {
				//	if a, e := msg, c.Expected[j]; a != e {
				//		log.Printf("%d, expected %v message, got %v", i, e, a)
				//	}
				//}
			}
		})
	})
})

type mockSQSClient struct {
	mock.Mock
}

func setSQSAttributes(queue string) {
	var name string
	var timeout int
	flag.StringVar(&name, "n", queue, "Queue name")
	flag.IntVar(&timeout, "t", 20, "(Optional) Timeout in seconds for long polling")
	flag.Parse()

	if len(name) == 0 {
		flag.PrintDefaults()
		exitErrorf("Queue name required")
	}
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
