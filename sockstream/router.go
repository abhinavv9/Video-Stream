package sockstream

import (
	"net/http"

	"log"
)

type Router struct {
	hub *Hub
}

func NewRouter(hub *Hub) *Router {
	return &Router{
		hub: hub,
	}
}

func (router *Router) Run() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveWs(router.hub, w, r)
	})
	log.Fatal(http.ListenAndServe(":5002", nil))
	router.hub.Run()
}
