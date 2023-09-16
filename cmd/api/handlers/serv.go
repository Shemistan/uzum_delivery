package handlers

import (
	"net/http"

	"github.com/Shemistan/uzum_delivery/internal/service"
)

type IHandlers interface {
	Healthz(w http.ResponseWriter, r *http.Request)
	GiveOrder(w http.ResponseWriter, r *http.Request)
	AddOrder(w http.ResponseWriter, r *http.Request)
}

func NewHandler(serv service.IService) IHandlers {
	return &handler{serv: serv}
}

type handler struct {
	serv service.IService
}
