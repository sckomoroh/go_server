package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type itemImpl struct {
	Id      string `json:"Id"`
	Message string `json:"Message"`
}

type itemsHandlerImpl struct {
	items []itemImpl
}

func (handler itemsHandlerImpl) GetHandlers() []ServerHandler {
	return []ServerHandler{
		{"/items/all", handler.GetAllItems},
		{"/items/{id}", handler.GetItemById},
	}
}
func (handler itemsHandlerImpl) GetItemById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	itemId := vars["id"]

	var targetItem *itemImpl = nil
	for _, item := range handler.items {
		if item.Id == itemId {
			targetItem = &item
		}
	}

	if targetItem != nil {
		json.NewEncoder(writer).Encode(&targetItem)
	} else {
		fmt.Fprintf(writer, "GetItemById result. Not found")
	}

	fmt.Println("Endpoint Hit: GetItemById")
}

func (handler itemsHandlerImpl) GetAllItems(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode(handler.items)
	fmt.Println("Endpoint Hit: GetAllItems")
}
