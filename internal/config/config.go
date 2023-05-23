package config

import (
	"flag"
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		App        App     `yaml:"app"`
		HTTP       HTTP    `yaml:"http"`
		Swagger    Swagger `yaml:"swagger"`
		PostgreSQL PostgreSQL
	}

	App struct {
		Debug bool `yaml:"debug"`
	}

	HTTP struct {
		PORT      string `yaml:"port"`
		HOST      string `yaml:"host"`
		PrefixAPI string `yaml:"prefixApi"`

		CORS struct {
			Debug              bool     `yaml:"debug"`
			AllowedMethods     []string `yaml:"allowed-methods"`
			AllowedOrigins     []string `yaml:"allowed-origins"`
			AllowCredentials   bool     `yaml:"allow-credentials"`
			AllowedHeaders     []string `yaml:"allowed-headers"`
			OptionsPassthrough bool     `yaml:"options-passthrough"`
			ExposedHeaders     []string `yaml:"exposed-headers"`
		} `yaml:"cors"`
	}

	Swagger struct {
		Path    string   `yaml:"path"`
		Title   string   `yaml:"title"`
		Version string   `yaml:"version"`
		Schemes []string `yaml:"schemes"`
	}

	PostgreSQL struct {
		Username string `env:"PSQL_USER" env-description:"username from database" env-required:"true"`
		Password string `env:"PSQL_PASSWORD" env-description:"password from database" env-required:"true"`
		Host     string `env:"PSQL_HOST" env-description:"host from database" env-required:"true"`
		Port     string `env:"PSQL_PORT" env-description:"port from database" env-required:"true"`
		Database string `env:"PSQL_DATABASE" env-description:"name database" env-required:"true"`
	}
)

const (
	pathConfigDefault  = "configs/config.local.yaml"
	FlagConfigPathName = "config"
)

var configPath string
var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		flag.StringVar(&configPath, FlagConfigPathName, pathConfigDefault, "this is app config file")
		flag.Parse()

		log.Print("config init")

		if err := godotenv.Load(); err != nil {
			log.Println("virtual environment file (.env) not found")
		}

		if configPath == "" {
			log.Fatal("config path is required")
		}

		instance = &Config{}

		if err := cleanenv.ReadConfig(configPath, instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Println(help)
			log.Fatal(err)
		}

	})

	return instance
}
