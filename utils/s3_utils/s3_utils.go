package s3_utils

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/cheggaaa/pb/v3"
	"github.com/sikalabs/slu/utils/vault_s3_utils"
)

func getAWSConfig(access_key, secret_key, region, endpoint string) (aws.Config, error) {
	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			access_key,
			secret_key,
			"",
		)),
	)
	if err != nil {
		return aws.Config{}, err
	}

	if region != "" {
		cfg.Region = region
	} else if endpoint != "" {
		cfg.Region = "us-east-1"
	}

	return cfg, nil
}

func getS3Client(cfg aws.Config, endpoint string) *s3.Client {
	if endpoint != "" {
		return s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.BaseEndpoint = aws.String(endpoint)
			o.UsePathStyle = true
		})
	}
	return s3.NewFromConfig(cfg)
}

func DeleteBucketWithObjects(
	access_key string,
	secret_key string,
	region string,
	endpoint string,
	bucket_name string,
) error {
	cfg, err := getAWSConfig(access_key, secret_key, region, endpoint)
	if err != nil {
		return err
	}

	client := getS3Client(cfg, endpoint)
	ctx := context.Background()

	// List and delete all objects
	paginator := s3.NewListObjectsV2Paginator(client, &s3.ListObjectsV2Input{
		Bucket:  aws.String(bucket_name),
		MaxKeys: aws.Int32(1000),
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			log.Println(err)
			return err
		}

		for _, obj := range page.Contents {
			_, err = client.DeleteObject(ctx, &s3.DeleteObjectInput{
				Bucket: aws.String(bucket_name),
				Key:    obj.Key,
			})
			if err != nil {
				log.Println(err)
			}
		}
	}

	// Delete the bucket
	_, err = client.DeleteBucket(ctx, &s3.DeleteBucketInput{
		Bucket: aws.String(bucket_name),
	})

	return err
}

func Upload(
	access_key string,
	secret_key string,
	region string,
	endpoint string,
	bucket_name string,
	key string,
	f io.ReadSeeker,
) error {
	return baseUpload(
		access_key,
		secret_key,
		region,
		endpoint,
		bucket_name,
		key,
		f,
		5,
		1,
	)
}

func DownloadToFile(
	access_key string,
	secret_key string,
	region string,
	endpoint string,
	bucket_name string,
	key string,
	localFilePath string,
) error {
	cfg, err := getAWSConfig(access_key, secret_key, region, endpoint)
	if err != nil {
		return err
	}

	client := getS3Client(cfg, endpoint)
	ctx := context.Background()

	file, err := os.Create(localFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	downloader := manager.NewDownloader(client)
	_, err = downloader.Download(ctx, file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket_name),
			Key:    aws.String(key),
		})
	if err != nil {
		return err
	}
	return nil
}

func baseUpload(
	access_key string,
	secret_key string,
	region string,
	endpoint string,
	bucket_name string,
	key string,
	f io.ReadSeeker,
	partSizeMB int,
	concurrency int,
) error {
	cfg, err := getAWSConfig(access_key, secret_key, region, endpoint)
	if err != nil {
		return err
	}

	client := getS3Client(cfg, endpoint)
	ctx := context.Background()

	uploader := manager.NewUploader(client, func(u *manager.Uploader) {
		u.PartSize = int64(partSizeMB) * 1024 * 1024 // The minimum/default allowed part size is 5MB
		u.Concurrency = concurrency                  // default is 5
	})

	size, _ := f.Seek(0, io.SeekEnd)
	f.Seek(0, io.SeekStart)

	bar := pb.Full.Start64(size)

	// create proxy reader
	barReader := bar.NewProxyReader(f)

	_, err = uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket_name),
		ACL:    types.ObjectCannedACLPrivate,
		Key:    aws.String(key),
		Body:   barReader,
	})
	if err != nil {
		return err
	}

	bar.Finish()

	return nil
}

func GetObjectPresignUrl(
	access_key string,
	secret_key string,
	region string,
	endpoint string,
	bucket_name string,
	key string,
	ttl time.Duration,
) (string, error) {
	cfg, err := getAWSConfig(access_key, secret_key, region, endpoint)
	if err != nil {
		return "", err
	}

	client := getS3Client(cfg, endpoint)
	ctx := context.Background()

	presignClient := s3.NewPresignClient(client)
	presignResult, err := presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket_name),
		Key:    aws.String(key),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = ttl
	})

	if err != nil {
		return "", err
	}

	return presignResult.URL, nil
}

func RemoveObjectsByAge(
	access_key string,
	secret_key string,
	region string,
	endpoint string,
	bucket_name string,
	age time.Duration,
) error {
	cfg, err := getAWSConfig(access_key, secret_key, region, endpoint)
	if err != nil {
		return err
	}

	client := getS3Client(cfg, endpoint)
	ctx := context.Background()

	paginator := s3.NewListObjectsV2Paginator(client, &s3.ListObjectsV2Input{
		Bucket:  aws.String(bucket_name),
		MaxKeys: aws.Int32(1000),
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			log.Println(err)
			return err
		}

		for _, obj := range page.Contents {
			if time.Since(*obj.LastModified) > age {
				fmt.Println("removing", *obj.Key)
				_, err = client.DeleteObject(ctx, &s3.DeleteObjectInput{
					Bucket: aws.String(bucket_name),
					Key:    obj.Key,
				})
				if err != nil {
					log.Println(err)
				}
			} else {
				fmt.Println("keeping", *obj.Key)
			}
		}
	}

	return nil
}

func GetS3SecretsFromVaultOrEnvOrDie(vaultPath string) (
	string, string, string, string, string,
) {
	accessKeyVault, secretKeyVault, regionVault,
		endpointVault, bucketNameVault, _ := vault_s3_utils.GetS3Secrets("secret/data/slu/upload")

	// Access Key
	var accessKey string
	accessKeyEnv := os.Getenv("SLU_UPLOAD_ACCESS_KEY")
	if accessKeyVault != "" {
		accessKey = accessKeyVault
	}
	if accessKeyEnv != "" {
		accessKey = accessKeyEnv
	}
	if accessKey == "" {
		log.Fatalln("SLU_UPLOAD_ACCESS_KEY is empty")
	}

	// Secret Key
	var secretKey string
	secretKeyEnv := os.Getenv("SLU_UPLOAD_SECRET_KEY")
	if secretKeyVault != "" {
		secretKey = secretKeyVault
	}
	if secretKeyEnv != "" {
		secretKey = secretKeyEnv
	}
	if secretKey == "" {
		log.Fatalln("SLU_UPLOAD_SECRET_KEY is empty")
	}

	// Region
	var region string
	regionEnv := os.Getenv("SLU_UPLOAD_REGION")
	if regionVault != "" {
		region = regionVault
	}
	if regionEnv != "" {
		region = regionEnv
	}

	// Endpoint
	var endpoint string
	endpointEnv := os.Getenv("SLU_UPLOAD_ENDPOINT")
	if endpointVault != "" {
		endpoint = endpointVault
	}
	if endpointEnv != "" {
		endpoint = endpointEnv
	}

	// Region, Endpoint Validation
	if region == "" && endpoint == "" {
		log.Fatalln("SLU_UPLOAD_REGION and SLU_UPLOAD_ENDPOINT are empty")
	}

	// Secret Key
	var bucketName string
	bucketNameEnv := os.Getenv("SLU_UPLOAD_BUCKET_NAME")
	if bucketNameVault != "" {
		bucketName = bucketNameVault
	}
	if bucketNameEnv != "" {
		bucketName = bucketNameEnv
	}
	if bucketName == "" {
		log.Fatalln("SLU_UPLOAD_BUCKET_NAME is empty")
	}

	return accessKey, secretKey, region, endpoint, bucketName
}
