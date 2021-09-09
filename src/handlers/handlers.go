package handlers

import (
	"net/http"
	"strconv"
)

type IHandlersProvider interface {
	GetHandlers() []ServerHandler
}

type ServerHandler struct {
	UrlPath string
	Handler func(writer http.ResponseWriter, request *http.Request)
}

func CreateHomePageHandler() IHandlersProvider {
	return homeHandlerImpl{}
}

func CreateItemsHandler() IHandlersProvider {
	itemsHandler := itemsHandlerImpl{}

	for i := 1; i < 10; i++ {
		itemsHandler.items = append(itemsHandler.items, itemImpl{
			strconv.FormatUint(uint64(i), 10),
			"Item" + strconv.FormatUint(uint64(i), 10),
		})
	}

	return &itemsHandler
}
