package api

import (
	"github.com/go-chi/chi"
	
	"demodesk/neko/internal/api/member"
	"demodesk/neko/internal/api/room"
	"demodesk/neko/internal/types"
)

type API struct {
	sessions   types.SessionManager
	remote     types.RemoteManager
	broadcast  types.BroadcastManager
	websocket  types.WebSocketHandler
}

func New(
	sessions types.SessionManager,
	remote types.RemoteManager,
	broadcast types.BroadcastManager,
	websocket types.WebSocketHandler,
) *API {
	// Init

	return &API{
		sessions:   sessions,
		remote:     remote,
		broadcast:  broadcast,
		websocket:  websocket,
	}
}

func (a *API) Mount(router *chi.Mux) {
	// all member routes
	memberHandler := member.New(a.sessions, a.websocket)
	router.Mount("/member", memberHandler.Router())

	// get room routes
	roomHandler := room.New(a.sessions, a.remote, a.broadcast, a.websocket)
	router.Mount("/room", roomHandler.Router())
}
