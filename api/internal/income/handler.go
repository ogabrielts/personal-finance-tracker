package income

import (
	"encoding/json"
	"expense-api/api/internal/utils"
	"net/http"
	"strconv"
)

type incomeHandler struct {
	service *incomeService
}

func newHandler(service *incomeService) *incomeHandler {
	return &incomeHandler{
		service: service,
	}
}

type incomeRequest struct {
	Name string `json:"name"`
	Value float32 `json:"value"`
	Category string `json:"category"`
}

func (ih *incomeHandler) CreateIncomeHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("token").(utils.Claims)
	if !ok {
		http.Error(w, "token failure", http.StatusInternalServerError)
		return
	}

	var incomeReq incomeRequest
	if err := json.NewDecoder(r.Body).Decode(&incomeReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	income, err := ih.service.CreateIncomeService(r.Context(), incomeReq.Name, incomeReq.Category, user.UserId, incomeReq.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(income)
}

func (ih *incomeHandler) GetIncomeByIdHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("token").(utils.Claims)
	if !ok {
		http.Error(w, "token failure", http.StatusInternalServerError)
		return
	}

	idStr := r.PathValue("id")
	incomeId, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	income, err := ih.service.GetIncomeByIdService(r.Context(), incomeId, user.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(income)
}

func (ih *incomeHandler) GetAllIncomesHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("token").(utils.Claims)
	if !ok {
		http.Error(w, "token failure", http.StatusInternalServerError)
		return
	}

	expenses, err := ih.service.GetAllIncomesService(r.Context(), user.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expenses)
}

func (ih *incomeHandler) UpdateIncomeHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("token").(utils.Claims)
	if !ok {
		http.Error(w, "token failure", http.StatusInternalServerError)
		return
	}

	idStr := r.PathValue("id")
	incomeId, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var incomeReq incomeRequest
	if err := json.NewDecoder(r.Body).Decode(&incomeReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	expense, err := ih.service.UpdateIncomeService(r.Context(), incomeId, user.UserId, incomeReq.Name, incomeReq.Category, incomeReq.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expense)
}

func (ih *incomeHandler) DeleteIncomeHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("token").(utils.Claims)
	if !ok {
		http.Error(w, "token failure", http.StatusInternalServerError)
		return
	}

	idStr := r.PathValue("id")
	incomeId, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := ih.service.DeleteIncomeService(r.Context(), incomeId, user.UserId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("income deleted")
}

func (ih *incomeHandler) FilterByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("token").(utils.Claims)
	if !ok {
		http.Error(w, "token failure", http.StatusInternalServerError)
		return
	}

	var incomeReq incomeRequest
	if err := json.NewDecoder(r.Body).Decode(&incomeReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	incomes, err := ih.service.FilterByCategoryService(r.Context(), incomeReq.Category, user.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(incomes)
}

func (ih *incomeHandler) FilterByDateHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("token").(utils.Claims)
	if !ok {
		http.Error(w, "token failure", http.StatusInternalServerError)
		return
	}

	type datesReq struct {
		DateFrom string `json:"datefrom"`
		DateTo string `json:"dateto"`
	}
	
	var dates datesReq

	if err := json.NewDecoder(r.Body).Decode(&dates); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	incomes, err := ih.service.FilterByDateService(r.Context(), dates.DateFrom, dates.DateTo, user.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(incomes)
}
