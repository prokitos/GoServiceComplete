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
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}

// все пути
func routers() *mux.Router {
	router := mux.NewRouter()

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	router.HandleFunc("/delete", deleteGetRequest).Methods(http.MethodDelete)
	router.HandleFunc("/insert", insertGetRequest).Methods(http.MethodPost)
	router.HandleFunc("/update", updateGetRequest).Methods(http.MethodPut)
	router.HandleFunc("/show", showsSpecGetRequest).Methods(http.MethodGet)

	router.HandleFunc("/singIn", signIn).Methods(http.MethodPost)
	router.HandleFunc("/getToken", getToken).Methods(http.MethodPost)
	router.HandleFunc("/useToken", GetAccName).Methods(http.MethodPost)
	router.HandleFunc("/reToken", RenewToken).Methods(http.MethodPost)

	router.HandleFunc("/test/{id}", showByIdRequest).Methods(http.MethodGet)
	router.HandleFunc("/test/{id}", delByIdRequest).Methods(http.MethodDelete)

	return router
}

// https://www.youtube.com/watch?v=DBZgt9iIWzk
// http://localhost:8889/swagger/index.html#/orders
// https://www.soberkoder.com/swagger-go-api-swaggo/
