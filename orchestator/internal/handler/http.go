package handler

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/JuanGQCadavid/arm7_cloudlet_facedetection_demo/orchestator/internal/domain"
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
	r.LoadHTMLGlob("internal/handler/templates/*")

	r.GET("/", hdl.GetAllImagesAnalyzed)

	r.GET("/upload", hdl.UploadView)
	r.POST("/upload", hdl.AnalyzeImage)
}

func (hdl *Handler) AnalyzeImage(c *gin.Context) {
	var (
		requestId         = uuid.NewString()
		logger            = log.With().Str("requestId", requestId).Logger()
		file, header, err = c.Request.FormFile("image")
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

	defer os.RemoveAll(filename)

	defer out.Close()
	if _, err = io.Copy(out, file); err != nil {
		logger.Error().Err(err).Msg("err while saving the image ")
		c.AbortWithError(500, err)
		return
	}

	result, err := hdl.srv.AnalyzeImage(logger.WithContext(c.Request.Context()), requestId, filename)

	if err != nil {
		logger.Error().Err(err).Msg("err while creating the image file")
		c.AbortWithError(500, err)
		return
	}

	// c.JSON(http.StatusOK, result)

	var (
		smallest = []*domain.ImageOuput{}
	)

	if len(result) > 1 {
		smallest = result[1:]
	}

	c.HTML(http.StatusOK, "result.tmpl", gin.H{
		"ObjectsDNS":     GetOutboundIP(),
		"OriginalSource": result[0],
		"SmallImages":    smallest,
	})
}

func (hdl *Handler) GetAllImagesAnalyzed(c *gin.Context) {

	c.HTML(http.StatusOK, "upload.tmpl", gin.H{
		"title": "Hello from Server-Side Rendering",
		"name":  "Gin User",
	})
}

func (hdl *Handler) UploadView(c *gin.Context) {

	c.HTML(http.StatusOK, "upload.tmpl", gin.H{
		"title": "Analyze",
		"path":  "/upload",
	})
}

// Get preferred outbound ip of this machine
func GetOutboundIP() string {

	// ips, err := net.LookupIP("host.docker.internal")

	// if err != nil {
	// 	log.Err(err).Msg("Ups")
	// } else {
	// 	for _, ip := range ips {
	// 		fmt.Println("Host IP:", ip)
	// 		return ip
	// 	}
	// }

	publicIp, oks := os.LookupEnv("public_ip")

	if oks {
		return publicIp
	}

	// We need to check this
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Err(err).Msg("Ups")
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}
