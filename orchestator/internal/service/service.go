package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/JuanGQCadavid/arm7_cloudlet_facedetection_demo/orchestator/internal/domain"
	"github.com/JuanGQCadavid/arm7_cloudlet_facedetection_demo/orchestator/internal/ports"
)

const (
	BASE_BUCKET = "images"
)

type OrchestatorService struct {
	metadataRepository ports.MetadataRepository
	objectRepository   ports.ObjectRepository
	mlService          ports.MlService
	baseBucket         string
}

func NewService(
	metadataRepository ports.MetadataRepository,
	objectRepository ports.ObjectRepository,
	mlService ports.MlService,
) *OrchestatorService {
	return &OrchestatorService{
		metadataRepository: metadataRepository,
		objectRepository:   objectRepository,
		mlService:          mlService,
		baseBucket:         BASE_BUCKET,
	}
}

func (srv *OrchestatorService) Init() error {
	if err := srv.objectRepository.CreateBucketIfNotExist(srv.baseBucket); err != nil {
		log.Err(err).Msg("Err while creating init bucket")
		return err
	}
	return nil
}

// Steps:

// 0. Generate an UUID for the image in the object storage
// 1. Save the image from the path into a object storage
// 2. Call ML with the path and the expected output folder
// 3. return the image list from ML.
func (srv *OrchestatorService) AnalyzeImage(ctx context.Context, requestId, imagePath string) ([]*domain.ImageOuput, error) {
	var (
		logger = log.Ctx(ctx)
		aux    = strings.Split(imagePath, ".")
		ext    = aux[len(aux)-1]
		// proccessID = uuid.NewString()
	)

	logger.Debug().Str("Image", imagePath).Msg("Start analyze")

	objectPath, err := srv.objectRepository.UploadFile(ctx, imagePath, srv.baseBucket, fmt.Sprintf("%s/original.%s", requestId, ext))

	if err != nil {
		logger.Err(err).Msg("Fail uploading the file")
		return nil, err
	}

	return srv.mlService.DetectFaces(ctx, objectPath, fmt.Sprintf("%s/%s", srv.baseBucket, requestId))
}
