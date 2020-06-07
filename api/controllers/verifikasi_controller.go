package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"rest/api/models"
	"rest/api/responses"
	"rest/api/utils/formaterror"
)

func (server *Server) UpdateVerifikasiUser(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	token := vars["token"]

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user := models.User{}
	updatedUser, err := user.UpdateVerifyUser(server.DB, token)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedUser)
}
