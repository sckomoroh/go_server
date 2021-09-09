package handlers

import (
	"net/http"
)

type IHomeHandler interface {
	IHandlersProvider

	HomePage(writer http.ResponseWriter, request *http.Request)
}
