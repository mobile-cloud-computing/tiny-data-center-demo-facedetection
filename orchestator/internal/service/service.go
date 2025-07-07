package service

import (
	"context"

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

func (srv *OrchestatorService) AnalyzeImage(ctx context.Context, requestId, imagePath string) ([]*domain.ImageOuput, error) {
	logger := log.Ctx(ctx)
	logger.Debug().Str("Image", imagePath).Msg("Start analyze")

	return nil, nil
}
