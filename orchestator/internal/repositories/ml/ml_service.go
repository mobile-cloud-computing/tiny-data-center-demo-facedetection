package ml

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/JuanGQCadavid/arm7_cloudlet_facedetection_demo/orchestator/internal/domain"
	"github.com/rs/zerolog/log"
)

type MLService struct {
	serviceEndpoint string
}

func NewMLService(serviceEndpoint string) *MLService {
	return &MLService{
		serviceEndpoint: serviceEndpoint,
	}
}

func (mls *MLService) DetectFaces(ctx context.Context, baseImagePath, outputPath string) ([]*domain.ImageOuput, error) {
	var (
		logger = log.Ctx(ctx)
		uri    = fmt.Sprintf("http://%s/ml/faces?base_image_path=%s&output_path=%s", mls.serviceEndpoint, baseImagePath, outputPath)
	)

	logger.Info().Str("Path", uri).Msg("Calling ML")

	req, err := http.NewRequest(http.MethodPost, uri, nil)

	if err != nil {
		logger.Err(err).Str("URI", uri).Str("baseImagePath", baseImagePath).Str("outputPath", outputPath).Msg("Cataplun creating request ML")
		return nil, err
	}

	response, err := http.DefaultClient.Do(req)

	if err != nil {
		logger.Err(err).Str("URI", uri).Str("baseImagePath", baseImagePath).Str("outputPath", outputPath).Msg("Cataplun calling ML")
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		logger.Err(err).Int("HttpStatus", response.StatusCode).Msg("Err while reading the body")
	}
	sb := string(body)

	var (
		result []*domain.ImageOuput
	)

	if err := json.Unmarshal(body, &result); err != nil {
		logger.Err(err).Int("HttpStatus", response.StatusCode).Msg("Error while unmarshaling the body")
	}

	logger.Info().Int("HttpStatus", response.StatusCode).Msg(sb)

	return result, nil
}
