package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	DatabaseUser       string `env:"DATABASE_USER"`
	DatabasePassword   string `env:"DATABASE_PASSWORD"`
	DatabasePort       string `env:"DATABASE_PORT"`
	DatabaseConnection string `env:"DATABASE_CONNECTION"`
	DatabaseHost       string `env:"DATABASE_HOST"`
	DatabaseDb         string `env:"DATABASE_DB"`
	DatabaseDriver     string `env:"DATABASE_DRIVER"`
	AppSecret          string `env:"APP_SECRET"`
	//StoragePath string `yaml:"storage_path" env-required:"true"`
	//HTTPServer  `yaml:"http_server"`
	//Clients     ClientsConfig `yaml:"clients"`
	//AppSecret   string        `yaml:"app_secret" env-required:"true" env:"APP_SECRET"`
}

//
//type HTTPServer struct {
//	Address     string        `yaml:"address" env-default:"localhost:8080"`
//	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
//	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
//	User        string        `yaml:"user" env-required:"true"`
//	Password    string        `yaml:"password" env-required:"true" env:"HTTP_SERVER_PASSWORD"`
//}
//
//type Client struct {
//	Address      string        `yaml:"address"`
//	Timeout      time.Duration `yaml:"timeout"`
//	RetriesCount int           `yaml:"retriesCount"`
//	Insecure     bool          `yaml:"insecure"`
//}
//
//type ClientsConfig struct {
//	SSO Client `yaml:"sso"`
//}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var config Config

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &config
}
