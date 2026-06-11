package handler

//Для каждого запроса создается http.Request с Body, указывающим на это соединение
//нужно закрыть соединение а то оно не вернется в пкл
import (
	"encoding/json"
	"extra_muse/internal/model"
	"extra_muse/internal/service"
	"net/http"
)

type BalanceHandler struct {

	BalanceService service.BalanceService
}

func (bh *BalanceHandler) BalanceAdd(w http.ResponseWriter, r *http.Request) {//всегда указатель, потому что запрос мутабелен, большой и должен быть единственным экземпляром на всю обработку.
	
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	var req struct{
		TgID int64
		Amount int64

	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "invalid JSON"}`, http.StatusBadGateway)
		return
	}

	if req.Amount <= 0 {
		http.Error(w, `{"error": "Amount must be more than 0"}`, http.StatusBadGateway)
		return
	}

	err := bh.BalanceService.AddBalance(
		model.NewBalanceAdd{
			Amount: req.Amount,
			TgID: req.TgID,
		},
	)

	if err != nil {
		http.Error(w, `{"error": "BalanceService.AddBalance"}`, http.StatusInternalServerError)
		return
	}


}