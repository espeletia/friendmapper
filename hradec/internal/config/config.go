package config

import "github.com/spf13/viper"

type Config struct {
	Server           ServerConfig
	Service          ServiceConfig
	DBConfig         DBConfig
	MigrationsConfig MigrationsConfig
	JWTConfig        JWTConfig
	HashConfig       HashConfig
	LoggingConfig    LoggingConfig
}

func LoadConfig() *Config {
	return &Config{
		Server:           loadServerConfig(),
		Service:          loadServiceConfig(),
		DBConfig:         loadDbConfig(),
		MigrationsConfig: loadMigrationsConfig(),
		JWTConfig:        loadJWTConfig(),
		HashConfig:       loadHashConfig(),
		LoggingConfig:    loadLoggingConfig(),
	}
}

func configViper(configName string) *viper.Viper {
	v := viper.New()
	v.AutomaticEnv()
	v.SetConfigName(configName)
	v.SetConfigType("yaml")
	v.AddConfigPath("./configurations/")
	return v
}
