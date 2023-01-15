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
	router.HandleFunc("/api/anime/{name}", handlers.DeleteAnime).Methods("DELETE")
	router.HandleFunc("/api/{name}", handlers.FindAnime).Methods("POST")
	router.HandleFunc("/api/update/{name}/{description}/{genre}", handlers.UpdateByTitle).Methods("POST")
	router.HandleFunc("/api/list/all", handlers.FindAll).Methods("GET")
	http.Handle("/", router)

	fmt.Printf("Starting server localhost%v", os.Getenv("PORT"))
	router.NotFoundHandler = router.NewRoute().BuildOnly().HandlerFunc(handlers.Err).GetHandler()
	http.ListenAndServe(os.Getenv("PORT"), router)
}

func SetHeaders(router *mux.Router) {
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			w.Header().Add("x-anime-token", os.Getenv("TOKEN"))
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
			next.ServeHTTP(w, r)
		})
	})
}
