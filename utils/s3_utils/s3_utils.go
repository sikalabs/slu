package s3_utils

import (
	"log"

	aws_aws "github.com/aws/aws-sdk-go/aws"
	aws_credentials "github.com/aws/aws-sdk-go/aws/credentials"
	aws_session "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	aws_s3 "github.com/aws/aws-sdk-go/service/s3"
)

func DeleteBucketWithObjects(
	access_key string,
	secret_key string,
	region string,
	endpoint string,
	bucket_name string,
) error {
	awsConfig := aws_aws.Config{
		Credentials: aws_credentials.NewStaticCredentials(
			access_key,
			secret_key,
			"",
		),
	}
	if region != "" {
		awsConfig.Region = aws_aws.String(region)
	}
	if endpoint != "" {
		awsConfig.Region = aws_aws.String(string("us-east-1"))
		awsConfig.S3ForcePathStyle = aws_aws.Bool(true)
		awsConfig.Endpoint = aws_aws.String(endpoint)
	}
	session, err := aws_session.NewSession(
		&awsConfig,
	)
	if err != nil {
		return err
	}

	svc := aws_s3.New(session)

	err = svc.ListObjectsPages(&aws_s3.ListObjectsInput{
		Bucket:  aws_aws.String(bucket_name),
		MaxKeys: aws_aws.Int64(1000),
	}, func(p *s3.ListObjectsOutput, last bool) (shouldContinue bool) {
		for _, obj := range p.Contents {
			_, err = svc.DeleteObject(&aws_s3.DeleteObjectInput{
				Bucket: aws_aws.String(bucket_name),
				Key:    obj.Key,
			})
			if err != nil {
				log.Println(err)
			}
		}
		return true
	})
	if err != nil {
		log.Println(err)
		return err
	}

	svc.DeleteBucket(&aws_s3.DeleteBucketInput{
		Bucket: aws_aws.String(bucket_name),
	})

	return nil
}
