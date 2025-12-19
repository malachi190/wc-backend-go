package config

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
)

type AppConfig struct {
	Name string `koanf:"name"`
	Port string `koanf:"port"`
	Env  string `koanf:"env"`
}

type ResendConfig struct {
	ApiKey string `koanf:"api_key"`
}

type JwtConfig struct {
	AccessSecret  string `koanf:"access_secret"`
	RefreshSecret string `koanf:"refresh_secret"`
}

type TmdbConfig struct {
	ApiAccessToken string `koanf:"api_access_token"`
	ApiKey         string `koanf:"api_key"`
	BaseUrl        string `koanf:"base_url"`
}

type FirebaseConfig struct {
	ProjectID   string `koanf:"project_id"`
	ClientEmail string `koanf:"client_email"`
	PrivateKey  string `koanf:"private_key"`
}

type Config struct {
	DatabaseUrl string         `koanf:"database_url"`
	App         AppConfig      `koanf:"app"`
	Resend      ResendConfig   `koanf:"resend"`
	Jwt         JwtConfig      `koanf:"jwt"`
	TMDB        TmdbConfig     `koanf:"tmdb"`
	Firebase    FirebaseConfig `koanf:"firebase"`
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	k := koanf.New(".")

	p := env.Provider("WATCHCIRCLE_", ".", func(s string) string {
		return strings.ToLower(strings.TrimPrefix(s, "WATCHCIRCLE_"))
	})

	if err := k.Load(p, nil); err != nil {
		return nil, fmt.Errorf("error loading env: %v", err)
	}

	config := &Config{}

	if err := k.Unmarshal("", config); err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}

	return config, nil
}
