package Services

import (
	"context"
	"ct-backend/Model"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"time"
)

type (
	IStorageService interface {
		UploadFile(object *Model.S3ObjectRequest) error
		GeneratePresignedURL(object *Model.S3UrlRequest) (string, error)
	}

	StorageService struct {
		svc *s3.Client
	}
)

func StorageServiceProvider(svc *s3.Client) *StorageService {
	return &StorageService{
		svc: svc,
	}
}

func (h *StorageService) UploadFile(object *Model.S3ObjectRequest) error {
	file, err := object.File.Open()
	if err != nil {
		return fmt.Errorf("unable to open file: %v", err)
	}
	defer file.Close()

	uploader := manager.NewUploader(h.svc)
	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:             aws.String(object.Bucket),
		Key:                aws.String(object.Key),
		Body:               file,
		ContentDisposition: aws.String("inline"),
	})

	return err
}

func (h *StorageService) GeneratePresignedURL(object *Model.S3UrlRequest) (string, error) {
	presignClient := s3.NewPresignClient(h.svc)

	req, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(object.Bucket),
		Key:    aws.String(object.Key),
	}, s3.WithPresignExpires(5*time.Minute))

	if err != nil {
		return "", err
	}

	return req.URL, nil
}
