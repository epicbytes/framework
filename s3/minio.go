package s3

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"log"
	"sync"
	"time"
)

type MinioStorage interface {
	GetBuckets(ctx context.Context) ([]minio.BucketInfo, error)
	GetLink(key string) string
	PutObject(ctx context.Context, key string, object io.Reader, length int64, contentType string) (minio.UploadInfo, error)
	FPutObject(ctx context.Context, key string, path string, contentType string) (minio.UploadInfo, error)
	GetObject(ctx context.Context, key string) ([]byte, error)
	ListObjects(ctx context.Context, prefix string) ([]*Object, error)
	RemoveObject(ctx context.Context, objectName string) error
}

type S3Option struct {
	Address   string `env:"ADDRESS"`
	AccessKey string `env:"ACCESS_KEY"`
	SecretKey string `env:"SECRET_KEY"`
	Bucket    string `env:"BUCKET"`
	Region    string `env:"REGION"`
}

type minioStorage struct {
	mu     sync.Mutex
	s3     *minio.Client
	config *S3Option
}

type Object struct {
	Key         string
	Size        int64
	ContentType string
	UpdatedAt   time.Time
}

func NewMinioStorage(config *S3Option) MinioStorage {
	minioClient, err := minio.New(config.Address, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.SecretKey, ""),
		Secure: false,
		Region: config.Region,
	})
	if err != nil {
		log.Fatalln(err)
	}

	return &minioStorage{
		config: config,
		s3:     minioClient,
	}
}

func (s *minioStorage) GetBuckets(ctx context.Context) ([]minio.BucketInfo, error) {
	buckets, err := s.s3.ListBuckets(ctx)
	if err != nil {
		return nil, err
	}
	return buckets, nil
}

func (s *minioStorage) GetLink(key string) string {
	if key == "" {
		return ""
	}
	sb := bytes.NewBufferString("https://")
	sb.WriteString(s.config.Address)
	sb.WriteString("/")
	sb.WriteString(s.config.Bucket)
	sb.WriteString("/")
	sb.WriteString(key)

	return sb.String()
}

func (s *minioStorage) PutObject(ctx context.Context, key string, object io.Reader, length int64, contentType string) (minio.UploadInfo, error) {
	return s.s3.PutObject(ctx, s.config.Bucket, key, object, length, minio.PutObjectOptions{ContentType: contentType})
}

func (s *minioStorage) FPutObject(ctx context.Context, key string, path string, contentType string) (minio.UploadInfo, error) {
	return s.s3.FPutObject(ctx, s.config.Bucket, key, path, minio.PutObjectOptions{ContentType: contentType})
}

func (s *minioStorage) GetObject(ctx context.Context, key string) ([]byte, error) {
	result, err := s.s3.GetObject(ctx, s.config.Bucket, key, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(result)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *minioStorage) ListObjects(ctx context.Context, prefix string) ([]*Object, error) {
	opts := minio.ListObjectsOptions{
		Recursive: true,
		Prefix:    prefix,
	}
	var objs []*Object
	for object := range s.s3.ListObjects(ctx, s.config.Bucket, opts) {
		if object.Err != nil {
			return nil, object.Err
		}
		objs = append(objs, &Object{
			Key:         object.Key,
			Size:        object.Size,
			ContentType: object.ContentType,
			UpdatedAt:   object.LastModified,
		})
	}
	return objs, nil
}

func (s *minioStorage) RemoveObject(ctx context.Context, objectName string) error {
	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
	}
	return s.s3.RemoveObject(ctx, s.config.Bucket, objectName, opts)
}
