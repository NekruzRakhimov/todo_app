package handler

import (
	_ "github.com/NekruzRakhimov/todo_app/docs"
	"github.com/NekruzRakhimov/todo_app/pkg/service"
	"github.com/nats-io/nats.go"
	"github.com/swaggo/http-swagger"
	"net/http"
)

type Handler struct {
	services *service.Service
	nats     *nats.Conn
}

func NewHandler(services *service.Service, nats *nats.Conn) *Handler {
	return &Handler{services: services, nats: nats}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	sm := http.NewServeMux()

	item := NewItem(h.services, h.nats)
	auth := NewAuth(h.services, h.nats)

	sm.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8880/swagger/doc.json"),
	))

	sm.HandleFunc("/auth/sign-up", auth.signUp)
	sm.HandleFunc("/auth/sign-in", auth.signIn)

	sm.Handle("/api/items", item.middleware(http.HandlerFunc(item.ItemsCR)))
	sm.Handle("/api/items/", item.middleware(item))
	sm.Handle("/api/items/bulk", item.middleware(http.HandlerFunc(item.bulkCreateItems)))

	return sm
}
