package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

const (
	configPrefix = "FIBER_PONGO2_PKGER"

	// EnvConfigPath represents the environment variable name which
	// should be read in case environment file needs to be read from some
	// user-defined location
	EnvConfigPath = "FIBER_PONGO2_PKGER_CONFIG_ENV_PATH"
)

// Config is the base struct which contains all the configuration for the
// application
type Config struct {

	// Address to run the webserver on, defaults to :3030
	Address string `envconfig:"ADDRESS" default:"0.0.0.0:3030"`
}

// LoadEnv will try to load .env file from the directory
// where it is currently running from, unless explicitly given
func LoadEnv() (Config, error) {

	pathToEnv := os.Getenv(EnvConfigPath)
	if pathToEnv == "" {
		pathToEnv = ".env"
	}

	log.Info().Msgf("Now reading config file %s", pathToEnv)
	var c Config
	err := godotenv.Load(pathToEnv)
	if err != nil {
		return c, err
	}

	err = envconfig.Process(configPrefix, &c)
	return c, err
}
