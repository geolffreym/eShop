package main

import (
	"github.com/gorilla/mux"
	"app/libs/confhandler"
	"app/controller/index"
	"app/controller/products"
	"net/http"
	"fmt"
)

type ConfigurationApp struct {
	Version string
	AppName string
	Http struct {
		Port string
	}
}

func main() {
	//Conf file
	var appConf ConfigurationApp
	confhandler.SetConf(&appConf, "app.json")

	//Router
	routes := mux.NewRouter()

	//Routes
	routes.HandleFunc("/", index.Index).Methods("GET")
	routes.HandleFunc("/products/", products.IndexProducts).
		Methods("GET").Queries("q", "{q}")

	routes.HandleFunc("/products/{id:[0-9A-Za-z]+}", products.IndexProductsByID).
		Methods("GET")

	//Listener
	fmt.Printf("Listening: 127.0.0.1%s \n", appConf.Http.Port)
	http.ListenAndServe(appConf.Http.Port, routes)

}
