package controllers

import (
	"net/http"

	"rest/api/responses"
	"github.com/gorilla/securecookie"
)


var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Golang Native by Adji Pangestu")

}
