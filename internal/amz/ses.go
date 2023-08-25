package amz

import (
	"api/internal/data"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

type SESClient struct {
	Client *sesv2.Client
}

func NewSesClient(ctx context.Context) *SESClient {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic(err)
	}
	return &SESClient{
		Client: sesv2.NewFromConfig(cfg),
	}
}

func (rcv SESClient) Send(ctx context.Context, e *data.Email) error {
	if e.Sender == "" {
		e.Sender = "team@freedomof.tech"
	}
	// Configure the email input
	emailInput := &sesv2.SendEmailInput{
		FromEmailAddress: aws.String(e.Sender),
		Destination: &types.Destination{
			ToAddresses: e.Recipients,
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Subject: &types.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(e.Subject),
				},
				Body: &types.Body{
					Text: &types.Content{
						Charset: aws.String("UTF-8"),
						Data:    aws.String(e.BodyText),
					},
				},
			},
		},
	}

	// Send the email
	_, err := rcv.Client.SendEmail(ctx, emailInput)
	return err
}
