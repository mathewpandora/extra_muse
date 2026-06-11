package handler

import (
	"encoding/json"
	"extra_muse/internal/model"
	"extra_muse/internal/service"
	"net/http"
)

type UserHandler struct{
	UserSerivce service.UserService
}

func (uh *UserHandler) Register(w http.ResponseWriter, r *http.Request)  {
    
	w.Header().Set("Content-Type", "application/json")

	var req struct {
        ID        int64  `json:"id"`
        Username  string `json:"username"`
        FirstName string `json:"first_name"`
        LastName  string `json:"last_name"`
    }

	//Создам декодер который умеет читать json из тела запроса

	//JsonDecoder := json.NewDecoder(r.Body)
	//JsonDecoder.Decode(&req) //декодируем и записываем в переменную
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "invalid JSON"}`, http.StatusBadRequest)
		return
	}

	

	//if req.PP - тут валидация

	 err := uh.UserSerivce.AddUser(

		model.NewUserData{
			TgID: req.ID,
			Username: req.Username,
    		FirstName: req.FirstName ,
    		LastName : req.LastName,
		},
	)

	if err != nil {
		http.Error(w, `{"error": "AddUser error"}`, http.StatusBadRequest)
		return
	}


}
