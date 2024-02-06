package app

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	_ "modular/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// запуск сервера
func MainServer() {

	router := routers()

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8888",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}

// все пути
func routers() *mux.Router {
	router := mux.NewRouter()

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	router.HandleFunc("/delete", deleteGetRequest).Methods(http.MethodDelete)
	router.HandleFunc("/insert", insertGetRequest).Methods(http.MethodPost)
	router.HandleFunc("/update", updateGetRequest).Methods(http.MethodPost)
	router.HandleFunc("/show", showsSpecGetRequest).Methods(http.MethodGet)

	return router
}

// https://www.youtube.com/watch?v=DBZgt9iIWzk
// http://localhost:8889/swagger/index.html#/orders
// https://www.soberkoder.com/swagger-go-api-swaggo/
