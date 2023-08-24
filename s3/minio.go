package s3

import (
	"bytes"
	"context"
	"github.com/epicbytes/framework/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog/log"
	"io"
	"net/url"
	"time"
)

type MinioStorageInt interface {
	GetBuckets(ctx context.Context) ([]minio.BucketInfo, error)
	GetLink(key string) string
	PresignedPutObject(ctx context.Context, key string, expires time.Duration) (*url.URL, error)
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

type MinioStorage struct {
	s3     *minio.Client
	Config *config.Config
	bucket string
}

type Object struct {
	Key         string
	Size        int64
	ContentType string
	UpdatedAt   time.Time
}

func (s *MinioStorage) Init(ctx context.Context) error {
	log.Debug().Msg("INITIAL S3")
	minioClient, err := minio.New(s.Config.S3.Address, &minio.Options{
		Creds:  credentials.NewStaticV4(s.Config.S3.AccessKey, s.Config.S3.SecretKey, ""),
		Secure: true,
		Region: s.Config.S3.Region,
	})
	if err != nil {
		log.Error().Err(err)
		return err
	}
	s.s3 = minioClient
	s.bucket = s.Config.S3.Bucket

	return nil
}
func (s *MinioStorage) Ping(context.Context) error {
	return nil
}
func (s *MinioStorage) Close() error {
	return nil
}

func (s *MinioStorage) GetBuckets(ctx context.Context) ([]minio.BucketInfo, error) {
	buckets, err := s.s3.ListBuckets(ctx)
	if err != nil {
		return nil, err
	}
	return buckets, nil
}

func (s *MinioStorage) GetLink(key string) string {
	if key == "" {
		return ""
	}
	sb := bytes.NewBufferString("https://")
	sb.WriteString(s.Config.S3.Address)
	sb.WriteString("/")
	sb.WriteString(s.bucket)
	sb.WriteString("/")
	sb.WriteString(key)

	return sb.String()
}

func (s *MinioStorage) PresignedPutObject(ctx context.Context, key string, expires time.Duration) (*url.URL, error) {
	return s.s3.PresignedPutObject(ctx, s.bucket, key, expires)
}

func (s *MinioStorage) PutObject(ctx context.Context, key string, object io.Reader, length int64, contentType string) (minio.UploadInfo, error) {
	return s.s3.PutObject(ctx, s.bucket, key, object, length, minio.PutObjectOptions{ContentType: contentType})
}

func (s *MinioStorage) FPutObject(ctx context.Context, key string, path string, contentType string) (minio.UploadInfo, error) {
	return s.s3.FPutObject(ctx, s.bucket, key, path, minio.PutObjectOptions{ContentType: contentType})
}

func (s *MinioStorage) GetObject(ctx context.Context, key string) ([]byte, error) {
	result, err := s.s3.GetObject(ctx, s.bucket, key, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(result)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *MinioStorage) ListObjects(ctx context.Context, prefix string) ([]*Object, error) {
	opts := minio.ListObjectsOptions{
		Recursive: true,
		Prefix:    prefix,
	}
	var objs []*Object
	for object := range s.s3.ListObjects(ctx, s.bucket, opts) {
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

func (s *MinioStorage) RemoveObject(ctx context.Context, objectName string) error {
	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
	}
	return s.s3.RemoveObject(ctx, s.bucket, objectName, opts)
}
