package router

import (
	"api/handlers"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func StartRouter() {
	router := mux.NewRouter()
	SetHeaders(router)

	router.HandleFunc("/", handlers.Ready)
	router.HandleFunc("/api/anime/{name}/{description}/{genre}", handlers.RegisterAnime).Methods("POST")
	router.HandleFunc("/api/update/{name}/{description}/{genre}", handlers.UpdateByTitle).Methods("POST")
	router.HandleFunc("/api/{name}", handlers.FindAnime).Methods("GET")
	router.HandleFunc("/api/list/all", handlers.FindAll).Methods("GET")
	router.HandleFunc("/api/anime/{name}", handlers.DeleteAnime).Methods("DELETE")
	http.Handle("/", router)

	fmt.Printf("Starting server localhost%v", Port())
	router.NotFoundHandler = router.NewRoute().BuildOnly().HandlerFunc(handlers.Err).GetHandler()
	http.ListenAndServe(Port(), router)
}
