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
	DB         DB         `yaml:"db"`
	HTTPServer HTTPServer `yaml:"http_server"`
	CorsOrigin []string   `yaml:"cors_origin" env-required:"true"`
	Cache      Cache      `yaml:"cache"`
	Auth       Auth       `yaml:"auth"`
}

type DB struct {
	Host     string `yaml:"host" env-required:"true"`
	Username string `yaml:"username" env-required:"true"`
	Port     string `yaml:"port" env-required:"true"`
	DBName   string `yaml:"dbname" env-required:"true"`
	SSLMode  string `yaml:"sslmode" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
}

type Cache struct {
	TTL time.Duration `yaml:"ttl" env-default:"5m"`
}

type Auth struct {
	AccessTokenTTL  time.Duration `yaml:"access_token_ttl" env-default:"15m"` // Длительность жизни аccessTokenTTL
	RefreshTokenTTL time.Duration `yaml:"refresh_token_ttl" env-default:"24h"`
	JWTSecret       string        `yaml:"jwt_secret" env-required:"true"`
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

	return cfg
}
