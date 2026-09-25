package reports

import (
	"encoding/json"
	"expense-api/internal/utils"
	"net/http"
)

type reportsHandler struct {
	service *reportsService
}

type reportRequest struct {
	Month string `json:"month"`
	Year int `json:"year"`
}

func newHandler(service *reportsService) *reportsHandler {
	return &reportsHandler{
		service: service,
	}
}

func (rh *reportsHandler) allReportHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("token").(utils.Claims)
	if !ok {
		http.Error(w, "token failure", http.StatusInternalServerError)
		return
	}

	report, err := rh.service.GetAllExpensesService(r.Context(), user.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(report)
}

func (rh *reportsHandler) monthlyReportHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("token").(utils.Claims)
	if !ok {
		http.Error(w, "token failure", http.StatusInternalServerError)
		return
	}

	var reportReq reportRequest
	if err := json.NewDecoder(r.Body).Decode(&reportReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	report, err := rh.service.MonthlyReportService(r.Context(), user.UserId, reportReq.Month, reportReq.Year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(report)
}
