package providers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

// CreateAWSConfig creates an AWS SDK v2 Config supporting custom endpoints (e.g., LocalStack) and default AWS.
func CreateAWSConfig(ctx context.Context, endpoint, region string) (aws.Config, error) {
	if region == "" {
		region = "us-east-1"
	}
	if endpoint != "" {
		return config.LoadDefaultConfig(ctx,
			config.WithRegion(region),
			config.WithBaseEndpoint(endpoint),
			config.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider("test", "test", ""),
			),
		)
	}
	return config.LoadDefaultConfig(ctx,
		config.WithRegion(region),
	)
}

// CreatAWSConfig is an alias for CreateAWSConfig to preserve backward compatibility with the original name.
func CreatAWSConfig(ctx context.Context, endpoint, region string) (aws.Config, error) {
	return CreateAWSConfig(ctx, endpoint, region)
}
