package s3

import (
	"context"
	"github.com/minio/minio-go/v7"
	"os"
	"strings"
	"time"
)

const (
	PathSeparator = "/"
)

// A HttpMinioObject implements http.File interface, returned by a S3
// Open method and can be served by the FileServer implementation.
type HttpMinioObject struct {
	Client *minio.Client
	Object *minio.Object
	Bucket string
	Prefix string
	IsDir  bool
}

func (h *HttpMinioObject) Close() error {
	return h.Object.Close()
}

func (h *HttpMinioObject) Read(p []byte) (n int, err error) {
	return h.Object.Read(p)
}

func (h *HttpMinioObject) Seek(offset int64, whence int) (int64, error) {
	return h.Object.Seek(offset, whence)
}

func (h *HttpMinioObject) Readdir(count int) ([]os.FileInfo, error) {
	// List 'N' number of objects from a Bucket-name with a matching Prefix.
	listObjectsN := func(bucket, prefix string, count int) (objsInfo []minio.ObjectInfo, err error) {
		i := 1
		for object := range h.Client.ListObjects(context.Background(), bucket, minio.ListObjectsOptions{
			Prefix:    prefix,
			Recursive: true,
		}) {
			if object.Err != nil {
				return nil, object.Err
			}
			i++
			// Verify if we have printed N objects.
			if i == count {
				return
			}
			objsInfo = append(objsInfo, object)
		}
		return objsInfo, nil
	}

	// List non-recursively first count entries for Prefix 'Prefix" Prefix.
	objsInfo, err := listObjectsN(h.Bucket, h.Prefix, count)
	if err != nil {
		return nil, os.ErrNotExist
	}
	var fileInfos []os.FileInfo
	for _, objInfo := range objsInfo {
		if strings.HasSuffix(objInfo.Key, PathSeparator) {
			fileInfos = append(fileInfos, objectInfo{
				ObjectInfo: minio.ObjectInfo{
					Key:          strings.TrimSuffix(objInfo.Key, PathSeparator),
					LastModified: objInfo.LastModified,
				},
				prefix: strings.TrimSuffix(objInfo.Key, PathSeparator),
				isDir:  true,
			})
			continue
		}
		fileInfos = append(fileInfos, objectInfo{
			ObjectInfo: objInfo,
		})
	}
	return fileInfos, nil
}

func (h *HttpMinioObject) Stat() (os.FileInfo, error) {
	if h.IsDir {
		return objectInfo{
			ObjectInfo: minio.ObjectInfo{
				Key:          h.Prefix,
				LastModified: time.Now().UTC(),
			},
			prefix: h.Prefix,
			isDir:  true,
		}, nil
	}

	objInfo, err := h.Object.Stat()
	if err != nil {
		return nil, os.ErrNotExist
	}

	return objectInfo{
		ObjectInfo: objInfo,
	}, nil
}
