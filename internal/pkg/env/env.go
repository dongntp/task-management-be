package env

import (
	"crypto/ecdsa"
	"os"
	"task-management-be/internal/pkg/logger"
	"task-management-be/internal/pkg/sensitive"
	"time"

	"github.com/caarlos0/env/v9"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

// HTTPLimit is limit of http request
type HTTPLimit struct {
	Timeout      time.Duration `yaml:"timeout"`
	MaxRetries   int           `yaml:"maxRetries"`
	MaxQueryRate int           `yaml:"maxQueryRate"`
}

type Environment struct {
	APIPort    int    `env:"API_PORT"                   envDefault:"3000"`
	BasePath   string `env:"BASE_PATH"                  envDefault:"v1"`
	ConfigFile string `env:"CONFIG_FILE"`
	// DB connection string
	DBConnectionString sensitive.Sensitive `env:"DB_CONNECTION_STRING"`
}

type Config struct {
	Environment
	HTTPLimit
	RewardsSetterPrivateKey *ecdsa.PrivateKey
}

const prefix = "0x"

func GetConfig() (cfg Config) {
	opts := env.Options{
		RequiredIfNoDef: true,
	}

	if err := env.ParseWithOptions(&cfg.Environment, opts); err != nil {
		logger.Logger.Fatal("unable to load environment variables", zap.Error(err))
	}

	configFile, err := os.Open(cfg.Environment.ConfigFile)
	if err != nil {
		logger.Logger.Fatal("unable to open config file", zap.Error(err))
	}
	defer configFile.Close()

	if err = yaml.NewDecoder(configFile).Decode(&cfg.HTTPLimit); err != nil {
		logger.Logger.Fatal("unable to decode config file", zap.Error(err))
	}

	return cfg
}
