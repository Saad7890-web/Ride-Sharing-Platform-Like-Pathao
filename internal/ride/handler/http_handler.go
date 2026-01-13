package handler

import (
	"encoding/json"
	"net/http"
	"ride/internal/ride/usecase"
)


type RideHandler struct {
	uc usecase.RideUsecase
}

func NewRideHandler(uc usecase.RideUsecase) *RideHandler {
	return &RideHandler{uc: uc}
}

type requestRideRequest struct {
	Fare int64 `json:"fare"`
}

func (h *RideHandler)RequestRide(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
	var req requestRideRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.Fare <= 0 {
		http.Error(w, "invalid fare", http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	ride, err := h.uc.RequestRide(r.Context(), userID, req.Fare)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ride)
}