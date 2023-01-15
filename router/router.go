package router

import (
	"api/handlers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func StartRouter() {
	router := mux.NewRouter()

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

	router.HandleFunc("/", handlers.Ready)
	router.HandleFunc("*", handlers.Err)
	router.HandleFunc("/anime/{name}/{description}/{genre}", handlers.RegisterAnime).Methods("POST")
	router.HandleFunc("/anime/{name}", handlers.DeleteAnime).Methods("DELETE")
	router.HandleFunc("/{name}", handlers.FindAnime).Methods("POST")
	router.HandleFunc("/update/{name}/{description}/{genre}", handlers.UpdateByTitle).Methods("POST")
	router.HandleFunc("/list/all", handlers.FindAll).Methods("GET")
	http.Handle("/", router)

	http.ListenAndServe(os.Getenv("PORT"), router)
}
