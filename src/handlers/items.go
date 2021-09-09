package handlers

import (
	"net/http"
)

type IItemsHandler interface {
	IHandlersProvider

	GetItemById(writer http.ResponseWriter, request *http.Request)
	GetAllItems(writer http.ResponseWriter, request *http.Request)
}
