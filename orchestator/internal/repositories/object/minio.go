package object

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog/log"
)

type MinioRepository struct {
	minioClient *minio.Client
}

func NewMinioRepository(endpoint, accessKeyID, secretAccessKey string) (*MinioRepository, error) {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})

	if err != nil {
		return nil, err
	}
	return &MinioRepository{
		minioClient: minioClient,
	}, nil
}

func (repo *MinioRepository) CreateBucketIfNotExist(bucket string) error {
	ctx := log.Logger.WithContext(context.Background())
	exist, err := repo.minioClient.BucketExists(ctx, bucket)

	if err != nil {
		log.Error().Err(err).Msg("err on CreateBucketIfNotExist - Checking")
		return err
	}

	if !exist {
		if err = repo.minioClient.MakeBucket(ctx, bucket, minio.MakeBucketOptions{}); err != nil {
			log.Error().Err(err).Msg("err on CreateBucketIfNotExist - Bucket creation")
			return err
		}
	}

	return nil
}

func (repo *MinioRepository) UploadFile(ctx context.Context, localPath, bucketName, objectKey string) (string, error) {
	var (
		logger = log.Ctx(ctx)
	)
	uploaInfo, err := repo.minioClient.FPutObject(ctx, bucketName, objectKey, localPath, minio.PutObjectOptions{})

	if err != nil {
		logger.Err(err).Msg("Err while colling Minioclient FPuntObject")
		return "", err
	}

	return fmt.Sprintf("%s/%s", uploaInfo.Bucket, uploaInfo.Key), nil
}
