package las

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

type Options struct {
	Region string `short:"r" env:"AWS_REGION" help:"The region to use."`
}

type Client struct {
	ses *sesv2.Client
}

func NewClient(opts *Options) (*Client, error) {
	optFns := []func(*config.LoadOptions) error{
		config.WithRetryer(func() aws.Retryer {
			return retry.AddWithMaxAttempts(retry.NewStandard(), 10)
		}),
	}

	if opts.Region != "" {
		optFns = append(optFns, config.WithRegion(opts.Region))
	}

	cfg, err := config.LoadDefaultConfig(context.Background(), optFns...)

	if err != nil {
		return nil, err
	}

	ses := sesv2.NewFromConfig(cfg)

	c := &Client{
		ses: ses,
	}

	return c, nil
}

func (c *Client) ListAddSuppressedDestinations(f func([]types.SuppressedDestinationSummary)) error {
	input := &sesv2.ListSuppressedDestinationsInput{PageSize: aws.Int32(1000)}

	for {
		output, err := c.ses.ListSuppressedDestinations(context.Background(), input)

		if err != nil {
			return err
		}

		f(output.SuppressedDestinationSummaries)

		if output.NextToken == nil {
			break
		}

		input.NextToken = output.NextToken
		time.Sleep(1 * time.Second)
	}

	return nil
}
