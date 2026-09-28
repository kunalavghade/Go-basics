package providers

import (
	"context"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	appConfig "github.com/kunalavghade/Go-basics/go-api/internal/config"
	"github.com/kunalavghade/Go-basics/go-api/internal/interfaces"
)

type S3Provider struct {
	client   *s3.Client
	uploader *manager.Uploader
	bucket   string
	endpoint string
}

func NewS3Provider(cfg *appConfig.Config) interfaces.UploadProvider {
	awsCfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(cfg.AWS.Region),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				cfg.AWS.AccessKeyID,
				cfg.AWS.SecretAccessKey,
				"",
			),
		),
	)
	if err != nil {
		panic("Failed To create S3 provider" + err.Error())
	}
	// config for localstack
	client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		if cfg.AWS.S3Endpoint != "" {
			o.BaseEndpoint = aws.String(cfg.AWS.S3Endpoint)
			o.UsePathStyle = true
		}
	})
	return &S3Provider{
		client:   client,
		uploader: manager.NewUploader(client),
		bucket:   cfg.AWS.S3Bucket,
		endpoint: cfg.AWS.S3Endpoint,
	}
}

func (s *S3Provider) UploadFile(file *multipart.FileHeader, path string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	res, err := s.uploader.Upload(
		context.Background(),
		&s3.PutObjectInput{
			Bucket: aws.String(s.bucket),
			Key:    aws.String(path),
			Body:   src,
		},
	)

	if err != nil {
		return "", err
	}
	return *res.Key, nil
}

func (s *S3Provider) DeleteFile(path string) error {
	_, err := s.client.DeleteObject(context.Background(), &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(path),
	})
	return err
}
