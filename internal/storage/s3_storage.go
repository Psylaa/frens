package storage

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/bwoff11/frens/internal/config"
	"github.com/google/uuid"
	"github.com/h2non/filetype"
)

type S3Storage struct {
	Bucket string
	Region string
}

func newS3Storage(cfg *config.StorageDetails) *S3Storage {
	return &S3Storage{
		Bucket: cfg.S3.Bucket,
		Region: cfg.S3.Region,
	}
}

func (ls *S3Storage) DetectFileType(data []byte) config.FileType {
	buf, _ := ioutil.ReadFile(string(data))
	kind, _ := filetype.Match(buf)

	switch kind.MIME.Type {
	case "image":
		return config.Image
	case "video":
		return config.Video
	case "audio":
		return config.Audio
	default:
		return config.Other
	}
}

func (s3s *S3Storage) SaveFile(id uuid.UUID, data []byte) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(s3s.Region),
	})

	if err != nil {
		return fmt.Errorf("failed to create AWS session: %w", err)
	}

	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s3s.Bucket),
		Key:    aws.String(id.String()),
		Body:   bytes.NewReader(data),
	})

	if err != nil {
		return fmt.Errorf("failed to upload to S3: %w", err)
	}

	return nil
}

func (s3s *S3Storage) LoadFile(id uuid.UUID) (io.ReadCloser, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(s3s.Region),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create AWS session: %w", err)
	}

	svc := s3.New(sess)

	result, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s3s.Bucket),
		Key:    aws.String(id.String()),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to fetch from S3: %w", err)
	}

	if result.Body == nil {
		return nil, errors.New("received empty body from S3")
	}

	return result.Body, nil
}

func (s3s *S3Storage) DeleteFile(id uuid.UUID) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(s3s.Region),
	})

	if err != nil {
		return fmt.Errorf("failed to create AWS session: %w", err)
	}

	svc := s3.New(sess)

	_, err = svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s3s.Bucket),
		Key:    aws.String(id.String()),
	})

	if err != nil {
		return fmt.Errorf("failed to delete from S3: %w", err)
	}

	return nil
}
