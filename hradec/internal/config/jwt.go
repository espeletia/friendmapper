package config

import "time"

type JWTConfig struct {
	Signature  string
	Expiration time.Duration
}

func loadJWTConfig() JWTConfig {
	JWTConfig := &JWTConfig{}
	v := configViper("jwt")
	v.BindEnv("JWT_SIGNATURE", "JWT_EXPIRATION")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	v.Unmarshal(JWTConfig)
	return *JWTConfig
}
