package ports

import "context"

type ObjectRepository interface {
	CreateBucketIfNotExist(bucket string) error
	UploadFile(ctx context.Context, localPath, bucketName, objectKey string) (string, error)
}
