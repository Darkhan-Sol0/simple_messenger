package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	webConfig struct {
		ServerConf struct {
			Address string `yaml:"server_address"`
		} `yaml:"server"`

		PgConf struct {
			Host     string `yaml:"pg_host"`
			Port     string `yaml:"pg_port"`
			Database string `yaml:"pg_database"`
			Username string `yaml:"pg_user"`
			Password string `yaml:"pg_password"`
		} `yaml:"postgresql"`

		RedisConf struct {
			Host     string `yaml:"redis_host"`
			Port     string `yaml:"redis_port"`
			Database string `yaml:"redis_database"`
			// Username string `yaml:"redis_user"`
			Password string `yaml:"redis_password"`
		} `yaml:"redis"`
	}

	Config interface {
		GetAddress() string

		GetPGHost() string
		GetPGPort() string
		GetPGName() string
		GetPGUser() string
		GetPGPassword() string

		GetRedisHost() string
		GetRedisPort() string
		GetRedisDB() string
		GetRedisPassword() string
	}
)

var (
	conf Config
	once sync.Once
)

func GetConfig() Config {
	once.Do(func() {
		conf = &webConfig{}
		path := "./config/config.yaml"
		if err := cleanenv.ReadConfig(path, conf); err != nil {
			log.Fatalf("error read config file %s: %v", path, err)
		}
	})
	return conf
}

func (c *webConfig) GetAddress() string {
	return c.ServerConf.Address
}

func (c *webConfig) GetPGHost() string {
	return c.PgConf.Host
}

func (c *webConfig) GetPGPort() string {
	return c.PgConf.Port
}

func (c *webConfig) GetPGName() string {
	return c.PgConf.Database
}

func (c *webConfig) GetPGUser() string {
	return c.PgConf.Username
}

func (c *webConfig) GetPGPassword() string {
	return c.PgConf.Password
}

func (c *webConfig) GetRedisHost() string {
	return ""
}

func (c *webConfig) GetRedisPort() string {
	return ""
}

func (c *webConfig) GetRedisDB() string {
	return ""
}

func (c *webConfig) GetRedisPassword() string {
	return ""
}
