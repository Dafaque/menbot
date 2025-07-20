package api

import (
	api "github.com/Dafaque/mentbot/internal/api/gen"
	"github.com/Dafaque/mentbot/internal/store"
)

type commandsUpdater interface {
	UpdateCommands()
}

type handler struct {
	store           store.Repository
	commandsUpdater commandsUpdater
}

func New(store store.Repository, commandsUpdater commandsUpdater) api.StrictServerInterface {
	return &handler{store: store, commandsUpdater: commandsUpdater}
}
