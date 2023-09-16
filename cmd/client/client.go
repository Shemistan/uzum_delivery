package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Shemistan/uzum_delivery/internal/models"
)

func main() {
	baseURL := "http://127.0.0.1:8080//deliver/v1/give-delivery/1234"
	client := http.Client{}

	r, err := client.Get(baseURL)
	if err != nil {
		log.Println(err.Error())
		return
	}

	defer r.Body.Close()
	//b, err := ioutil.ReadAll(r.Body)
	b, err := io.ReadAll(r.Body)
	//_, err := r.Body.Read(b)
	defer r.Body.Close()
	if err != nil {
		log.Println(err.Error())

	}

	order := &models.Order{}

	err = json.Unmarshal(b, &order)
	if err != nil {
		log.Println(err.Error())
		return
	}

	fmt.Print(order)
}
