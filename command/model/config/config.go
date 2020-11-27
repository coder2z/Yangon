package config

import (
	"fmt"
	"github.com/spf13/viper"
	"yangon/pkg/database"
)

const (
	// DefaultConfigurationName is the default name of configuration
	defaultConfigurationName = "config"

	// DefaultConfigurationPath the default location of the configuration file
	defaultConfigurationPath = "./config"
)

type Cfg struct {
	Mysql *database.Options
}

func cfg() *Cfg {
	return &Cfg{
		Mysql: database.NewDatabaseOptions(),
	}
}

func TryLoadFromDisk() (*Cfg, error) {
	viper.SetConfigName(defaultConfigurationName)
	viper.AddConfigPath(defaultConfigurationPath)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, err
		} else {
			return nil, fmt.Errorf("error parsing configuration file %s", err)
		}
	}
	conf := cfg()
	if err := viper.Unmarshal(conf); err != nil {
		return nil, err
	}
	return conf, nil
}
