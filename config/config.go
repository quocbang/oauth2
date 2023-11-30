package config

import (
	"flag"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Oauth2Config struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	RedirectURL  string `yaml:"call_back_url"`
}

type Oauth2 struct {
	Google Oauth2Config `yaml:"google"`
	Github Oauth2Config `yaml:"github"`
}

type PostgresConfig struct {
	Name     string `yaml:"name"`
	Address  string `yaml:"address"`
	Port     int    `yaml:"port"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Schema   string `yaml:"schema"`
}

type DatabaseGroup struct {
	Postgres PostgresConfig `yaml:"postgres"`
}

type InternalAuthInfo struct {
	SecretKey            string `yaml:"secret_key"`
	AccessTokenDuration  int    `yaml:"access_token_duration"`
	RefreshTokenDuration int    `yaml:"refresh_token_duration"`
}

var c Config

type Config struct {
	DevMode      bool             `yaml:"dev_mode"`
	Oauth2       Oauth2           `yaml:"oauth2"`
	InternalAuth InternalAuthInfo `yaml:"internal_auth"`
	Database     DatabaseGroup    `yaml:"database"`
	MigratePath  string           `yaml:"migrate_path"`
}

func init() {
	configPath := flag.String("config-path", "", "")
	// submit parse flag panic if run without flag
	flag.Parse()

	data, err := os.ReadFile(*configPath)
	if err != nil {
		log.Fatalf("failed to read config file, error: %v", err)
	}

	if err := yaml.Unmarshal(data, &c); err != nil {
		log.Fatalf("failed to unmarshal config data, error: %v", err)
	}
}

func GetConfig() Config {
	return c
}
