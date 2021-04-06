package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

// StartServer - Start API
func StartServer() {
	log.Info("📡 API Server starting to listening on port 8080")

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/spaceship", spaceshipCreateHandler).Methods("POST")
	r.HandleFunc("/spaceship/{id}", spaceshipGetHandler).Methods("GET")
	r.HandleFunc("/spaceship/{id}", spaceshipUpdateHandler).Methods("PUT")
	r.HandleFunc("/spaceship/{id}", spaceshipDeleteHandler).Methods("DELETE")
	r.HandleFunc("/spaceship/list", spaceshipListHandler).Methods("POST")

	handler := cors.Default().Handler(r)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
