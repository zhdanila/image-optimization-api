package app

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
)

type Config struct {
	Env      string `mapstructure:"ENV" validate:"required"`
	HTTPPort string `mapstructure:"HTTP_PORT" validate:"required"`

	RabbitMQHost string `mapstructure:"RABBIT_MQ_HOST" validate:"required"`

	S3Region    string `mapstructure:"S3_REGION" validate:"required"`
	S3AccessKey string `mapstructure:"S3_ACCESS_KEY" validate:"required"`
	S3SecretKey string `mapstructure:"S3_SECRET_KEY" validate:"required"`
	S3Bucket    string `mapstructure:"S3_BUCKET" validate:"required"`
}

func (cnf *Config) GetPort() string {
	return cnf.HTTPPort
}

func NewConfig() (*Config, error) {
	var err error

	baseCnf, err := New()
	if err != nil {
		return nil, err
	}

	if err = baseCnf.Load(baseCnf); err != nil {
		return nil, err
	}

	return baseCnf, nil
}

func New() (*Config, error) {
	cnf := Config{
		Env: os.Getenv(`ENV`),
	}

	return &cnf, nil
}

func (cnf *Config) Load(extConf any) error {
	var err error

	viper.SetConfigFile(`.env`)
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %v", err)
	}

	if err = viper.Unmarshal(cnf); err != nil {
		return fmt.Errorf("configuration 1 can't be loaded: %s", err)
	}
	if err = viper.Unmarshal(extConf); err != nil {
		return fmt.Errorf("configuration 2 can't be loaded: %s", err)
	}

	zap.L().Info(fmt.Sprintf("Loaded config for %s environment.\n", cnf.Env))

	return nil
}
