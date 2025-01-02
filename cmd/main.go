package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/Dafaque/mentbot/internal/bot/tg"
	"github.com/Dafaque/mentbot/internal/config"
	"github.com/Dafaque/mentbot/internal/handlers/handler"
	"github.com/Dafaque/mentbot/internal/store"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		flag.Usage()
		log.Fatal(err)
	}

	db, err := store.New(cfg.DBPath)
	if err != nil {
		log.Fatal(err)
	}

	h := handler.New(db, cfg)

	b := tg.NewBot(cfg.BotToken, h)
	if err != nil {
		log.Fatal(err)
	}

	b.Start()
	defer b.Stop()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
