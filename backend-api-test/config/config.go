package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"time"
)

// App config struct
type Config struct {
	Server  ServerConfig
	Cookie  Cookie
	Session Session
	Logger  Logger
	JWT     JWT
}

// @desc: Server config struct
type ServerConfig struct {
	AppVersion        string
	Port              string
	PprofPort         string
	Mode              string
	CookieName        string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	SSL               bool
	CtxDefaultTimeout time.Duration
	CSRF              bool
	Debug             bool
	MaxConnectionIdle time.Duration
	Timeout           time.Duration
	MaxConnectionAge  time.Duration
	Time              time.Duration
	ShutdownTimeout   time.Duration
	Url               string
	NameService       string
}

// Logger config
type Logger struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

// Cookie config
type Cookie struct {
	Name     string
	MaxAge   int
	Secure   bool
	HTTPOnly bool
}

// Session config
type Session struct {
	Prefix string
	Name   string
	Expire int
}

type JWT struct {
	SecretKey    string
	AccessToken  string
	RefreshToken string
}

// Load config file from given path
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

// Parse config file
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}

// Get config
func GetConfig(configPath string) (*Config, error) {
	cfgFile, err := LoadConfig(configPath)
	if err != nil {
		return nil, err
	}

	cfg, err := ParseConfig(cfgFile)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func GetConfigPath(configPath string) string {
	fmt.Println("configPath: ", configPath)
	if configPath == "production" {
		return "./config/config-production"
	} else if configPath == "staging" {
		return "./config/config-staging"
	} else if configPath == "develop" {
		return "./config/config-develop"
	} else {
		return "./config/config-local"
	}
}
