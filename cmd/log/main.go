package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/tcotav/logtest/data"
	"github.com/tcotav/logtest/middleware"
)

const (
	// ServiceName is the name of the app
	ServiceName    = "logtest"
	ServiceVersion = "0.1.0"
)

func initViper() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil { // exit on failed config file read
		log.Panic().Err(err).Msg("Error while reading config file")
	}
}

func main() {

	// set up viper config
	initViper()
	port := viper.GetInt("port")
	loglevel := viper.GetString("loglevel")

	// set up zerolog log level -- removed a few levels
	switch loglevel {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	// set up gin
	ginMode := "release"
	gin.SetMode(ginMode) // need to set this to turn off default DEBUG noise from gin logging
	router := gin.New()
	// set our custom request logger here w/gin
	router.Use(middleware.RequestLogger())
	// set up otelgin middleware
	router.Use(gin.Recovery())

	log.Debug().Str("App", ServiceName).Str("Call", "main").Msg("Testing debug and loglevel")

	// hack in some dummy handler
	router.GET("/ping", func(c *gin.Context) {
		s, err := data.GetThing()
		if err != nil {
			log.Error().Err(err).Str("App", ServiceName).Str("Call", "ping").Msg("Error in main getting thing")
		}
		log.Info().Str("thing", s).Str("App", ServiceName).Str("Call", "ping").Msg("Got thing")

		data.TimeSomething()

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Info().Str("App", ServiceName).Str("Call", "main").Msgf("Starting server on port %d", port)
	router.Run(fmt.Sprintf(":%d", port))
}
