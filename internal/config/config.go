package config

import "github.com/daddydemir/RAJ/internal/infrastructure"

type Config struct {
	DatabaseUrl string
	RedisUrl    string
}

func LoadConfig() *Config {
	client := infrastructure.NewVaultClient()

	secrets := infrastructure.GetSecrets(client)

	return &Config{
		DatabaseUrl: secrets["database_url"].(string),
		RedisUrl:    secrets["redis_url"].(string),
	}
}
