package main

import (
	"flag"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/JuanGQCadavid/arm7_cloudlet_facedetection_demo/orchestator/internal/handler"
	"github.com/JuanGQCadavid/arm7_cloudlet_facedetection_demo/orchestator/internal/ports"
	"github.com/JuanGQCadavid/arm7_cloudlet_facedetection_demo/orchestator/internal/repositories/object"
	"github.com/JuanGQCadavid/arm7_cloudlet_facedetection_demo/orchestator/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

const (
	OBJECT_ACCES_KEY_ENV_NAME        = "object_access_key"
	OBJECT_SECRET_ACCES_KEY_ENV_NAME = "object_secret_access_key"
	OBJECT_DNS_ENV_NAME              = "object_dns"
)

var (
	// Interfaces
	objectRepository ports.ObjectRepository
)

func init() {
	var (
		err error
	)
	// Object
	access, okey := os.LookupEnv(OBJECT_ACCES_KEY_ENV_NAME)
	if !okey {
		log.Fatal().Str("env-missing", OBJECT_ACCES_KEY_ENV_NAME).Msg("Missing env variable")
	}
	secret, okey := os.LookupEnv(OBJECT_SECRET_ACCES_KEY_ENV_NAME)
	if !okey {
		log.Fatal().Str("env-missing", OBJECT_SECRET_ACCES_KEY_ENV_NAME).Msg("Missing env variable")
	}
	dns, okey := os.LookupEnv(OBJECT_DNS_ENV_NAME)
	if !okey {
		log.Fatal().Str("env-missing", OBJECT_DNS_ENV_NAME).Msg("Missing env variable")
	}

	objectRepository, err = object.NewMinioRepository(dns, access, secret)
	if err != nil {
		panic(err.Error())
	}

}

func main() {

	var (
		router = gin.Default()
		srv    = service.NewService(nil, objectRepository, nil)
		hdl    = handler.NewHandler(srv)
		debug  = flag.Bool("debug", false, "sets log level to debug")
	)

	flag.Parse()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	if err := srv.Init(); err != nil {
		log.Fatal().Err(err).Msg("We could not init the service")
	}

	hdl.SetRouter(router)
	router.Run()
}
