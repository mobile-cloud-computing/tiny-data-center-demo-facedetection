package ports

import (
	"context"

	"github.com/JuanGQCadavid/arm7_cloudlet_facedetection_demo/orchestator/internal/domain"
)

type MlService interface {
	DetectFaces(ctx context.Context, baseImagePath, outputPath string) ([]*domain.ImageOuput, error)
}
