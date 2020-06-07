package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"rest/api/auth"
	"rest/api/utils"
	"rest/api/models"
	"rest/api/responses"
	"rest/api/utils/formaterror"

	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}

	if token == "false" {
		responses.JSON(w, http.StatusUnprocessableEntity, map[string]string{
			"error": "Email belum terferifikasi",
		})
		return
	}

	responses.JSON(w, http.StatusOK, map[string]string{
		"token": token,
	})
}

func (server *Server) Register(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userCreated, err := user.SaveUser(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	token := auth.GenerateToken(userCreated.Email)
	user.UpdateToken(server.DB, token, uint32(userCreated.ID))

	url_verification := os.Getenv("URL_CLIENT") + "/auth/verifikasi/" + userCreated.TokenVerification

	bodyEmail := "Verifikasi akun anda dengan klik link berikut ini <br><br> <a href='" + url_verification + "'>"+url_verification+"</a>"

	utils.SendMail(userCreated.Email, bodyEmail)

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	responses.JSON(w, http.StatusCreated, userCreated)
}

func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	user := models.User{}

	err = server.DB.Debug().Table("users").Where("email = ?", email).Take(&user).Error

	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	if user.EmailVerification == false {
		return "false", err
	}
	return auth.CreateToken(user.ID)
}

func (server *Server) GetToken(w http.ResponseWriter, r *http.Request) {

	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	user := models.User{}
	userGotten, err := user.TokenMe(server.DB, uint32(tokenID))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, userGotten)
}
