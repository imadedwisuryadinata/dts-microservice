package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/context"
	"github.com/imadedwisuryadinata/dts-microservice/utils"
	"github.com/imadedwisuryadinta/dts-microservice/menu-service/config"
	"github.com/imadedwisuryadinta/dts-microservice/menu-service/entity"
)

type AuthHandler struct {
	Config config.Auth
}

//menjalankan validasi terlebih dahulu => nexthandler
func (handler *AuthHandler) ValidateAdmin(nextHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := http.NewRequest(http.MethodPost, handler.Config.Host+"/auth/validate", nil)
		if err != nil {
			utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		request.Header = r.Header
		authResponse, err := http.DefaultClient.Do(request)
		if err != nil {
			utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
			return
		}
		defer authResponse.Body.Close()
		body, err := ioutil.ReadAll(authResponse.Body)
		if err != nil {
			utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		var authResult entity.AuthResponse
		err = json.Unmarshal(body, &authResult)

		if authResponse.StatusCode != 200 {

			utils.WrapAPIError(w, r, authResult.ErrorDetails, authResponse.StatusCode)
			return
		}

		context.Set(r, "user", authResult.Data.Username)
		nextHandler(w, r)

		if authResponse.StatusCode != http.StatusOK {
			utils.WrapAPIError(w, r, "invalid auth", authResponse.StatusCode)
			return
		}
		nextHandler(w, r)
	}
}
