package env

import (
	"os"
	"task-management-be/internal/pkg/logger"
	"task-management-be/internal/pkg/sensitive"
	"time"

	"github.com/caarlos0/env/v9"
	"github.com/getkin/kin-openapi/openapi3"
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
	APIPort         int    `env:"API_PORT"                   envDefault:"3000"`
	BasePath        string `env:"BASE_PATH"                  envDefault:"v1"`
	ConfigFile      string `env:"CONFIG_FILE"`
	OpenAPIFilePath string `env:"OPEN_API_FILE_PATH"`
	// DB connection string
	DBConnectionString sensitive.Sensitive `env:"DB_CONNECTION_STRING"`
	AdminAccessToken   sensitive.Sensitive `env:"ADMIN_ACCESS_TOKEN"`
}

type Config struct {
	Environment
	HTTPLimit
	OpenAPIPaths *openapi3.Paths
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
	cfg.OpenAPIPaths = loadPathsFromDocs(cfg.Environment.OpenAPIFilePath)

	return cfg
}

func loadPathsFromDocs(openAPIFile string) *openapi3.Paths {
	data, err := os.ReadFile(openAPIFile) // Change to "openapi.json" if needed
	if err != nil {
		logger.Logger.Fatal("Failed to read OpenAPI spec", zap.Error(err))
	}

	// Parse the OpenAPI spec
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromData(data)
	if err != nil {
		logger.Logger.Fatal("Failed to parse OpenAPI spec", zap.Error(err))
	}
	return doc.Paths
}
