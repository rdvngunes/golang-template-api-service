package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Configuration struct {
	App           App           `yaml:"App"`
	Database      Database      `yaml:"Database"`
	Auth          Auth          `yaml:"Auth"`
	Services      Services      `yaml:"Services"`
	AmazonStorage AmazonStorage `yaml:"AmazonStorage"`
}

type Database struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	DBName   string `yaml:"DBName"`
}

type AmazonStorage struct {
	AccessKey         string   `yaml:"AccessKey"`
	SecretKey         string   `yaml:"SecretKey"`
	BucketUrl         string   `yaml:"BucketUrl"`
	BucketName        string   `yaml:"BucketName"`
	Region            string   `yaml:"Region"`
	ObjectKey         string   `yaml:"ObjectKey"`
	AllowedExtensions []string `yaml:"AllowedExtensions"`
	FileSizeLimit     int      `yaml:"FileSizeLimit"`
}

type Services struct {
	TestServiceUrl string `yaml:"TestServiceUrl"`
}

type App struct {
	Port int `yaml:"Port"`
}

type Auth struct {
	SecretKeyPath string `yaml:"SecretKeyPath"`
}

func LoadViperConfig() *Configuration {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev" // Default to dev environment
	}

	var configuration Configuration
	//Work well with docker file also any change here please make change in Docker also
	viper.SetConfigFile(fmt.Sprintf("./app/config/settings/config-%s.yaml", env))
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error while reading the config file: %v, config file path: %v", err, viper.ConfigFileUsed())
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct: %v, config file path: %v", err, viper.ConfigFileUsed())
	}

	return &configuration
}
