package handler

import (
	"net/http"

	"github.com/imadedwisuryadinata/dts-microservice/utils"
	"github.com/imadedwisuryadinta/dts-microservice/menu-service/config"
)

type AuthHandler struct {
	Config config.Auth
}

//menjalankan validasi terlebih dahulu => nexthandler
func (handler *AuthHandler) ValidateAdmin(nextHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := http.NewRequest(http.MethodPost, handler.Config.Host+"/validate-admin", nil)
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

		// responseBody, err := ioutil.ReadAll(authResponse.Body)
		// if err != nil {
		// 	utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		// var resposeData map[string]interface{}
		// err = json.Unmarshal(responseBody, &responseData)
		// if err != nil {
		// 	utils.WrapAPIError(w, r, "invalid auth", authRespose.StatusCode)
		// 	return
		// }

		if authResponse.StatusCode != http.StatusOK {
			utils.WrapAPIError(w, r, "invalid auth", authResponse.StatusCode)
			return
		}
		nextHandler(w, r)
	}
}
