package ports

import "github.com/JuanGQCadavid/arm7_cloudlet_facedetection_demo/orchestator/internal/domain"

type MlService interface {
	DetectFaces(baseImagePath, outputPath string) (domain.ImageOuput, error)
}
