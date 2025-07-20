package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/Dafaque/mentbot/assets"
	api "github.com/Dafaque/mentbot/internal/api/gen"
	"github.com/Dafaque/mentbot/internal/bot/tg"
	"github.com/Dafaque/mentbot/internal/config"
	apiHandler "github.com/Dafaque/mentbot/internal/handlers/api"
	handler "github.com/Dafaque/mentbot/internal/handlers/tg"
	"github.com/Dafaque/mentbot/internal/store"
	"github.com/Dafaque/mentbot/internal/webhandler"
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

	go server(db)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	b.Stop()
	log.Println("close db error:", db.Done())

}

func server(db store.Repository) error {
	fs := webhandler.New(assets.Web, "web", "index.html")
	ssh := api.NewStrictHandler(apiHandler.New(db), []api.StrictMiddlewareFunc{})
	mux := http.NewServeMux()
	mux.Handle("/", fs)

	return http.ListenAndServe(":8080", api.HandlerFromMuxWithBaseURL(ssh, mux, "/api"))
}
