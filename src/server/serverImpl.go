package server

import (
	"log"
	"net/http"

	"handlers"

	"github.com/gorilla/mux"
)

func CreateServer() serverImpl {
	var server = serverImpl{}
	server.providers = append(
		server.providers,
		handlers.CreateHomePageHandler(),
		handlers.CreateItemsHandler())

	return server
}

type serverImpl struct {
	providers []handlers.IHandlersProvider
}

func (server serverImpl) Start() {
	log.Println("Start server")
	router := mux.NewRouter().StrictSlash(true)

	for _, providerItem := range server.providers {
		for _, handlerItem := range providerItem.GetHandlers() {
			router.HandleFunc(handlerItem.UrlPath, handlerItem.Handler)
		}
	}

	log.Fatal(http.ListenAndServe(":10000", router))
}

func (server serverImpl) Stop() {
	// TODO:
}
