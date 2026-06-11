package handler

import (
	"encoding/json"
	"extra_muse/internal/service"
	"net/http"
)

type GenerationHandler struct{
	GenerationService service.GenerationService
}

func (gh *GenerationHandler) Generate(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()
	var req struct{

	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, `{"error": "invalid JSON"}`, http.StatusBadRequest)
				return
	}
}