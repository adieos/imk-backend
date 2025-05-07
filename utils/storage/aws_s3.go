package storage

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"strings"

	"github.com/adieos/imk-backend/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type (
	AwsS3 interface {
		UploadFile(filename string, file *multipart.FileHeader, folderName string, mv ...string) (string, error)
		UpdateFile(objectKey string, f *multipart.FileHeader, mv ...string) (string, error)
		DeleteFile(objectKey string) error
		GetPublicLink(objectKey string) string
		GetObjectKeyFromLink(link string) string
	}

	awsS3 struct {
		client *s3.Client
		bucket string
		region string
	}
)

func NewAwsS3() AwsS3 {
	bucket := os.Getenv("S3_BUCKET")
	region := os.Getenv("AWS_REGION")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			os.Getenv("AWS_ACCESS_KEY"),
			os.Getenv("AWS_SECRET_KEY"),
			"",
		)),
	)
	if err != nil {
		panic(fmt.Sprintf("failed to load AWS configuration: %v", err))
	}

	client := s3.NewFromConfig(cfg)

	return &awsS3{
		client: client,
		bucket: bucket,
		region: region,
	}
}

// Gunakan untuk mengupload file ke s3 dimana defaultnya mengizinkan semua jenis mimetypa
func (a *awsS3) UploadFile(filename string, f *multipart.FileHeader, folderName string, mv ...string) (string, error) {
	file, err := f.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	mimetype, err := utils.GetMimetype(file)
	if err != nil {
		return "", err
	}

	if len(mv) > 0 {
		flag := false
		for _, m := range mv {
			if mimetype == m {
				flag = true
				break
			}
		}

		if !flag {
			return "", fmt.Errorf("invalid mimetype")
		}
	}

	objectKey := fmt.Sprintf("%s/%s", folderName, filename)

	_, err = a.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(a.bucket),
		Key:         aws.String(objectKey),
		Body:        file,
		ContentType: aws.String(mimetype),
	})
	if err != nil {
		return "", err
	}

	return objectKey, nil
}

// Gunakan untuk mengupdate file ke s3 dimana defaultnya mengizinkan semua jenis mimetypa
func (a *awsS3) UpdateFile(objectKey string, f *multipart.FileHeader, mv ...string) (string, error) {
	file, err := f.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	mimetype, err := utils.GetMimetype(file)
	if err != nil {
		return "", err
	}

	if len(mv) > 0 {
		flag := false
		for _, m := range mv {
			if mimetype == m {
				flag = true
				break
			}
		}

		if !flag {
			return "", fmt.Errorf("invalid mimetype")
		}
	}

	_, err = a.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(a.bucket),
		Key:         aws.String(objectKey),
		Body:        file,
		ContentType: aws.String(mimetype),
	})
	if err != nil {
		return "", err
	}

	return objectKey, nil
}

// Gunakan untuk menghapus file ke s3 dimana defaultnya mengizinkan semua jenis mimetypa
func (a *awsS3) DeleteFile(objectKey string) error {
	_, err := a.client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(a.bucket),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return err
	}
	return nil
}

func (a *awsS3) GetObjectKeyFromLink(link string) string {
	pref := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/", a.bucket, a.region)

	if !strings.HasPrefix(link, pref) {
		return ""
	}

	objectKey := strings.TrimPrefix(link, pref)
	return objectKey
}

func (a *awsS3) GetPublicLink(objectKey string) string {
	publicURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", a.bucket, a.region, objectKey)
	return publicURL
}
