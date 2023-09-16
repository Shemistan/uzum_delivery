package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Shemistan/uzum_delivery/internal/models"
)

func (h *handler) Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *handler) AddOrder(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("Body.Read", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	coordinate := &models.Coordinate{}
	order := &models.Order{
		Coordinate: coordinate,
	}

	err = json.Unmarshal(b, &order)
	if err != nil {
		log.Println("Unmarshal", err.Error())

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println(fmt.Printf("%+v\n", order)) // Записали в БД

	w.WriteHeader(http.StatusCreated)

}
