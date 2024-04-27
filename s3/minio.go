package s3

import (
	"bytes"
	"context"
	"crypto/tls"
	"github.com/epicbytes/framework/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog/log"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"
)

type MinioFS struct {
}

type MinioStorageInt interface {
	GetBuckets(ctx context.Context) ([]minio.BucketInfo, error)
	GetLink(key string) string
	PresignedPutObject(ctx context.Context, key string, expires time.Duration) (*url.URL, error)
	PutObject(ctx context.Context, key string, object io.Reader, length int64, contentType string) (minio.UploadInfo, error)
	FPutObject(ctx context.Context, key string, path string, contentType string) (minio.UploadInfo, error)
	GetObject(ctx context.Context, key string) ([]byte, error)
	ListObjects(ctx context.Context, prefix string) ([]*Object, error)
	RemoveObject(ctx context.Context, objectName string) error
	Init(ctx context.Context) error
	Ping(context.Context) error
	Open(name string) (http.File, error)
	Close() error
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
	u, err := url.Parse(s.Config.S3.Address)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	tlsConfig := &tls.Config{}
	tlsConfig.InsecureSkipVerify = u.Scheme == "https"

	var transport http.RoundTripper = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       tlsConfig,
	}
	minioClient, err := minio.New(u.Host, &minio.Options{
		Creds:        credentials.NewStaticV4(s.Config.S3.AccessKey, s.Config.S3.SecretKey, ""),
		Secure:       u.Scheme == "https",
		Region:       s.Config.S3.Region,
		BucketLookup: minio.BucketLookupAuto,
		Transport:    transport,
	})

	if err != nil {
		log.Error().Err(err).Send()
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
func (s *MinioStorage) GetClient() *minio.Client {
	return s.s3
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
	sb := bytes.NewBufferString("")
	if s.Config.S3.Secure {
		sb.WriteString("https://")
	} else {
		sb.WriteString("http://")
	}
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

// Open - implements http.Filesystem implementation.
func (s *MinioStorage) Open(name string) (http.File, error) {
	name = path.Join("/", name)
	if name == PathSeparator || pathIsDir(context.Background(), s, name) {
		return &HttpMinioObject{
			Client: s.GetClient(),
			Object: nil,
			IsDir:  true,
			Bucket: s.bucket,
			Prefix: strings.TrimSuffix(name, PathSeparator),
		}, nil
	}

	name = strings.TrimPrefix(name, PathSeparator)
	obj, err := getObject(context.Background(), s, name)
	if err != nil {
		return nil, os.ErrNotExist
	}
	file := &HttpMinioObject{
		Client: s.GetClient(),
		Object: obj,
		IsDir:  false,
		Bucket: s.bucket,
		Prefix: name,
	}
	return file, nil
}
