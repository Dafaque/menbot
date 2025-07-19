package api

import (
	api "github.com/Dafaque/mentbot/internal/api/gen"
	"github.com/Dafaque/mentbot/internal/store"
)

type handler struct {
	store store.Repository
}

func New(store store.Repository) api.StrictServerInterface {
	return &handler{store: store}
}
