package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/imadedwisuryadinta/dts-microservice/auth-service/database"
	"github.com/imadedwisuryadinta/dts-microservice/auth-service/utils"
	"gorm.io/gorm"
)

type Auth struct {
	Db *gorm.DB
}

func (db *Auth) ValidateAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	authToken := r.Header.Get("Authorization")
	res, err := database.Validate(authToken, db.Db)
	if err != nil {
		utils.WrapAPIError(w, r, err.Error(), http.StatusForbidden)
		return
	}

	utils.WrapAPIData(w, r, database.Auth{
		Username: res.Username,
		Token:    res.Token,
	}, http.StatusOK, "success")
	return
}

func (db *Auth) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	//TODO untuk signup
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		utils.WrapAPIError(w, r, "can't read body", http.StatusBadRequest)
		return
	}

	var signup database.Auth

	err = json.Unmarshal(body, &signup)
	if err != nil {
		utils.WrapAPIError(w, r, "error unmarshal"+err.Error(), http.StatusBadRequest)
		return
	}

	signup.Token = utils.IdGenerator()
	err = signup.SignUp(db.Db)
	if err != nil {
		utils.WrapAPIError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	utils.WrapAPISuccess(w, r, "Success", http.StatusOK)

}

func (db *Auth) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		utils.WrapAPIError(w, r, "can't read body", http.StatusBadRequest)
		return
	}

	var login database.Auth

	err = json.Unmarshal(body, &login)
	if err != nil {
		utils.WrapAPIError(w, r, "error unmarshal"+err.Error(), http.StatusBadRequest)
		return
	}

	res, err := login.Login(db.Db)
	if err != nil {
		utils.WrapAPIError(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	utils.WrapAPIData(w, r, database.Auth{
		Username: res.Username,
		Token:    res.Token,
	}, http.StatusOK, "success")

}
