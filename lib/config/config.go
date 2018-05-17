package config

import (
	"os"
)

type Config struct {
	DB      DB
	Session Session
	Github  Github
}

type DB struct {
	User string
	Pass string
	Host string
	Port string
}

type Session struct {
	Name   string
	Secret string
}

type Github struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Scopes       []string
}

func New() *Config {
	return &Config{
		DB: DB{
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
		},
		Session: Session{
			Name:   os.Getenv("SESS_NAME"),
			Secret: os.Getenv("SESS_SECRET"),
		},
		Github: Github{
			ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
			ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
			RedirectURL:  os.Getenv("GITHUB_OAUTH_CALLBACK"),
			Scopes:       []string{"read:org"},
		},
	}
}
