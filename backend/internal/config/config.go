package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env string `yaml:"env" env-default:"local"`
	//StoragePath string     `yaml:"storage_path" env-required:"true"`
	DB         DB
	HTTPServer HTTPServer `yaml:"http_server"`
	CorsOrigin []string   `yaml:"cors_origin" env-required:"true"`
	Cache      Cache      `yaml:"cache"`
	Auth       Auth       `yaml:"auth"`
}

type DB struct {
	Host     string `env:"POSTGRES_HOST" env-required:"true"`
	Username string `env:"POSTGRES_USERNAME" env-required:"true"`
	Port     string `env:"POSTGRES_PORT" env-required:"true"`
	DBName   string `env:"POSTGRES_DBNAME" env-required:"true"`
	SSLMode  string `env:"POSTGRES_SSLMODE" env-required:"true"`
	Password string `env:"POSTGRES_PASSWORD" env-required:"true"`
}

type Cache struct {
	TTL time.Duration `yaml:"ttl" env-default:"5m"`
}

type Auth struct {
	AccessTokenTTL  time.Duration `yaml:"access_token_ttl" env-default:"15m"` // Длительность жизни аccessTokenTTL
	RefreshTokenTTL time.Duration `yaml:"refresh_token_ttl" env-default:"24h"`
	JWTSecret       string        `env:"AUTH_JWT_SECRET" env-required:"true"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"8082"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"30s"`
	CertFile    string        `yaml:"cert_file" env-required:"true"` // Путь к SSL-сертификату
	KeyFile     string        `yaml:"key_file" env-required:"true"`  // Путь к приватному ключу
}

func MustLoad() Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatalf("CONFIG_PATH is not set")
	}
	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("config file %s does not exist: %v", configPath, err)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %v", err)
	}
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("cannot read config: %v", err)
	}

	return cfg
}
