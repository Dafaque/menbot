package config

import (
	"errors"
	"flag"
	"os"
	"path/filepath"
)

type Config struct {
	BotToken string
	DBPath   string
}

var (
	flagBotToken = flag.String("bot", "", "Bot token")
	flagDBPath   = flag.String("db", "storage", "DB path")
)

func NewConfig() (*Config, error) {
	flag.Parse()
	if *flagBotToken == "" {
		return nil, errors.New("bot token is required")
	}
	if _, err := os.Stat(*flagDBPath); err != nil {
		err = os.MkdirAll(filepath.Dir(*flagDBPath), 0755)
		if err != nil {
			return nil, err
		}
	}
	cfg := &Config{
		BotToken: *flagBotToken,
		DBPath:   *flagDBPath,
	}
	return cfg, nil
}
