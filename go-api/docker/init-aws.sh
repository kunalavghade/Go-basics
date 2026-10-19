#!/bin/bash
echo "Initializing LocalStack AWS resources..."

# Create S3 Bucket
awslocal s3 mb s3://ecommerse-upload || true

# Create SQS Queue
awslocal sqs create-queue --queue-name task-queue || true

echo "LocalStack initialization complete!"
