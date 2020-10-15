package handler

func AddMenu(w http.ResponseWriter, r http.Request) {
	utils.WrapAPISuccess(w, r, "success", http.StatusOk)
}
