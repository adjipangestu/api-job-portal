package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"rest/api/models"
	"rest/api/responses"
	"rest/api/utils/formaterror"
)

func (server *Server) UpdateVerifikasiUser(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	token := vars["token"]
	body, _ := ioutil.ReadAll(r.Body)
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	code := keyVal["code"]

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user := models.User{}
	updatedUser, err := user.UpdateVerifyUser(server.DB, token, code)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	fmt.Println(updatedUser)
	msg := models.Msg{
		Status:  200,
		Message: "success",
	}

	responses.JSON(w, http.StatusOK, msg)
}

func (server *Server) ForgotPassword(w http.ResponseWriter, r *http.Request)  {
	var err error
	vars := mux.Vars(r)
	token := vars["token"]
	body, _ := ioutil.ReadAll(r.Body)
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	code := keyVal["code"]
}
