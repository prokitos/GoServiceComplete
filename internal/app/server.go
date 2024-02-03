package app

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	_ "modular/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func MainServer() {

	router := mux.NewRouter()

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	router.HandleFunc("/delete", deleteGetRequest).Methods(http.MethodDelete)
	router.HandleFunc("/insert", insertGetRequest).Methods(http.MethodPost)
	router.HandleFunc("/update", updateGetRequest).Methods(http.MethodPost)
	router.HandleFunc("/show", showsSpecGetRequest).Methods(http.MethodGet)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8889",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}

// https://www.youtube.com/watch?v=DBZgt9iIWzk
// http://localhost:8889/swagger/index.html#/orders
// https://www.soberkoder.com/swagger-go-api-swaggo/
