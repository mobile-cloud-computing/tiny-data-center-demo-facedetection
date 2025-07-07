package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/JuanGQCadavid/arm7_cloudlet_facedetection_demo/orchestator/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	srv *service.OrchestatorService
}

func NewHandler(srv *service.OrchestatorService) *Handler {
	return &Handler{
		srv: srv,
	}
}

func (hdl *Handler) SetRouter(r *gin.Engine) {
	r.GET("/", hdl.GetAllImagesAnalyzed)
	r.POST("/", hdl.AnalyzeImage)
}

func (hdl *Handler) AnalyzeImage(c *gin.Context) {
	var (
		requestId         = uuid.NewString()
		logger            = log.With().Str("requestId", requestId).Logger()
		file, header, err = c.Request.FormFile("upload")
	)

	if err != nil {
		logger.Error().Err(err).Msg("err while fetching the image from the request ")
		c.AbortWithError(500, err)
		return
	}

	logger.Info().Msg("Hello world")
	ext := strings.Split(header.Filename, ".")[1]
	filename := fmt.Sprintf("./tmp/%s.%s", requestId, ext)
	out, err := os.Create(filename)

	if err != nil {
		logger.Error().Err(err).Msg("err while creating the image file")
		c.AbortWithError(500, err)
		return
	}

	defer out.Close()
	if _, err = io.Copy(out, file); err != nil {
		logger.Error().Err(err).Msg("err while saving the image ")
		c.AbortWithError(500, err)
		return
	}

	hdl.srv.AnalyzeImage(logger.WithContext(c.Request.Context()), requestId, filename)

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (hdl *Handler) GetAllImagesAnalyzed(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
