package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/verma-kartik/marketplace-api/internal/models"
	"github.com/verma-kartik/marketplace-api/internal/services"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

type Handler struct {
	Router  *mux.Router
	Server  *http.Server
	Service *services.ProductService
}

func NewHandler(service *services.ProductService) *Handler {
	h := &Handler{
		Service: service,
	}
	h.Router = mux.NewRouter()

	h.mapRoutes()
	h.Router.Use(JSONMiddleware)
	h.Server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: h.Router,
	}

	return h
}

func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world!")
	})

	h.Router.HandleFunc("/api/v1/product", h.PostProduct).Methods("POST")
	h.Router.HandleFunc("/api/v1/product/{serialNumber}", h.GetProduct).Methods("GET")
}

func (h *Handler) PostProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "inside the post product handler")
	var prod []models.Product
	fmt.Fprintf(w, "created an empty product")
	err := json.NewDecoder(r.Body).Decode(&prod)
	fmt.Fprintf(w, "trying the body product")
	if err != nil {
		fmt.Fprintf(w, "could not decode the product")
		log.Print(err)
		return
	}

	err = h.Service.CreateProduct(r.Context(), &prod[0])
	if err != nil {
		log.Print(err)
		return
	}

	if err := json.NewEncoder(w).Encode(prod); err != nil {
		panic(err)
	}
}

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serialnumber := vars["serialNumber"]
	if serialnumber == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sn, err := strconv.ParseInt(serialnumber, 10, 32)
	if err != nil {
		panic(err)
	}

	p := models.Product{}

	p, err = h.Service.GetProduct(r.Context(), int32(sn))
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(p); err != nil {
		panic(err)
	}

}

func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15+time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)

	log.Println("shutdown gracefully")

	return nil
}
