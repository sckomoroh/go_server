package handlers

import (
	"fmt"
	"net/http"
)

type homeHandlerImpl struct{}

func (handler homeHandlerImpl) GetHandlers() []ServerHandler {
	return []ServerHandler{{"/home/", handler.HomePage}}
}

func (handler homeHandlerImpl) HomePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: HomePage")
}
