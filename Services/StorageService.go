package Services

import (
	"context"
	"ct-backend/Config"
	"ct-backend/Model"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/minio/minio-go/v7"
	"log"
	"time"
)

type (
	IStorageService interface {
		UploadFile(object *Model.S3ObjectRequest) error
		GeneratePresignedURL(object *Model.S3UrlRequest) (string, error)
	}

	S3Service struct {
		svc *s3.Client
	}

	MinioService struct {
		minio *minio.Client
	}
)

func S3ServiceProvider() *S3Service {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-southeast-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	svc := s3.NewFromConfig(cfg)

	return &S3Service{
		svc: svc,
	}
}

func MinioServiceProvider() *MinioService {
	minioClient := Config.SetupMinioConnection()

	log.Printf("%#v\n", minioClient)

	return &MinioService{
		minio: minioClient,
	}
}

func (m MinioService) UploadFile(object *Model.S3ObjectRequest) error {
	file, err := object.File.Open()
	if err != nil {
		return fmt.Errorf("unable to open file: %v", err)
	}
	defer file.Close()

	info, err := m.minio.PutObject(context.TODO(), object.Bucket, object.Key, file, object.File.Size, minio.PutObjectOptions{
		ContentType: object.File.Header.Get("Content-Type"),
	})

	if err != nil {
		return fmt.Errorf("unable to upload file: %v", err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", object.Key, info.Size)

	return err
}

func (m MinioService) GeneratePresignedURL(object *Model.S3UrlRequest) (string, error) {
	presignedURL, err := m.minio.PresignedGetObject(
		context.TODO(),
		object.Bucket,
		object.Key,
		5*time.Minute,
		nil)

	if err != nil {
		return "", err
	}

	return presignedURL.String(), nil
}

func (h *S3Service) UploadFile(object *Model.S3ObjectRequest) error {
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

func (h *S3Service) GeneratePresignedURL(object *Model.S3UrlRequest) (string, error) {
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
