package config

import (
	"github.com/rs/zerolog"
)

func init() {
	// set the global log level to info
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}
