package main

import (
	"github.com/gin-gonic/gin"

	"github.com/rs/zerolog/log"
	"github.com/tcotav/logtest/data"
	"github.com/tcotav/logtest/middleware"
)

const (
	// AppName is the name of the app
	AppName = "logtest"
)

func main() {
	ginMode := "release"
	gin.SetMode(ginMode) // need to set this to turn off default DEBUG noise from gin logging
	router := gin.New()
	// set our custom request logger here w/gin
	router.Use(middleware.RequestLogger())
	router.Use(gin.Recovery())

	// hack in some dummy handler
	router.GET("/ping", func(c *gin.Context) {
		s, err := data.GetThing()
		if err != nil {
			log.Error().Err(err).Str("App", AppName).Str("Call", "main").Msg("Error in main getting thing")
		}
		log.Info().Str("thing", s).Str("App", AppName).Str("Call", "main").Msg("Got thing")

		data.TimeSomething()

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Info().Str("App", AppName).Str("Call", "main").Msg("Starting server on port 8080")
	router.Run(":8080")
}
