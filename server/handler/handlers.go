package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"petproj/model"
	"petproj/server/service"
)

func AlgoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req model.TreeTarget
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Ошибка декодирования json на серваке", http.StatusBadRequest)
		return
	}
	resp := service.Helper(req)

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println("Ошибка ответа сервера")
		return
	}
}
