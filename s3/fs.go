package s3

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
)

func pathIsDir(ctx context.Context, s3 *MinioStorage, name string) bool {
	name = strings.Trim(name, PathSeparator) + PathSeparator
	listCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	objCh := s3.GetClient().ListObjects(listCtx,
		s3.bucket,
		minio.ListObjectsOptions{
			Prefix:  name,
			MaxKeys: 1,
		})
	for range objCh {
		cancel()
		return true
	}
	return false
}

func getObject(ctx context.Context, s3 *MinioStorage, name string) (*minio.Object, error) {
	names := []string{name, name + "/index.html", name + "/index.htm"}
	names = append(names, "/404.html")
	for _, n := range names {
		obj, err := s3.GetClient().GetObject(ctx, s3.bucket, n, minio.GetObjectOptions{})
		if err != nil {
			log.Info().Err(err).Send()
			continue
		}

		_, err = obj.Stat()
		if err != nil {
			// do not log "file" in bucket not found errors
			if minio.ToErrorResponse(err).Code != "NoSuchKey" {
				log.Info().Err(err).Send()
			}
			continue
		}

		return obj, nil
	}

	return nil, os.ErrNotExist
}
