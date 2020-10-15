package handler

import (
	"net/http"

	"github.com/imadedwisuryadinata/dts-microservice/utils"
)

func AddMenu(w http.ResponseWriter, r *http.Request) {
	utils.WrapAPISuccess(w, r, "success", http.StatusOK)
}
