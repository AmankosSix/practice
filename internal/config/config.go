package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
	"time"
)

const (
	defaultHttpPort               = "8000"
	defaultHttpRWTimeout          = 10 * time.Second
	defaultHttpMaxHeaderMegabytes = 1

	EnvLocal = "local"
)

const (
	UsersTable = "users"
)

type (
	Config struct {
		Environment string
		HTTP        HTTPConfig
		Postgres    Postgres
	}

	HTTPConfig struct {
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
	}

	Postgres struct {
		User     string
		Password string
		Host     string
		Port     string
		Name     string `mapstructure="databaseName"`
		SSLMode  string `mapstructure="sslmode"`
	}
)

func Init(configsDir string) (*Config, error) {
	populateDefaults()
	fmt.Println(godotenv.Load())
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	if err := parseConfigFile(configsDir, os.Getenv("APP_ENV")); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)

	return &cfg, nil
}

func setFromEnv(cfg *Config) {
	cfg.Postgres.User = os.Getenv("POSTGRES_USER")
	cfg.Postgres.Password = os.Getenv("POSTGRES_PASS")
	cfg.Postgres.Host = os.Getenv("POSTGRES_HOST")
	cfg.Postgres.Port = os.Getenv("POSTGRES_PORT")
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("postgres.databaseName", &cfg.Postgres.Name); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("postgres.sslmode", &cfg.Postgres.SSLMode); err != nil {
		return err
	}

	return viper.UnmarshalKey("http", &cfg.HTTP)
}

func parseConfigFile(folder, env string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if env == EnvLocal {
		return nil
	}

	viper.SetConfigName(env)

	return viper.MergeInConfig()
}

func populateDefaults() {
	viper.SetDefault("http.port", defaultHttpPort)
	viper.SetDefault("http.max_header_megabytes", defaultHttpMaxHeaderMegabytes)
	viper.SetDefault("http.timeouts.read", defaultHttpRWTimeout)
	viper.SetDefault("http.timeouts.write", defaultHttpRWTimeout)
}
