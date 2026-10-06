package expense

import (
	"encoding/json"
	"errors"
	"expense-api/api/internal/utils"
	"net/http"
	"strconv"
)

type expenseHandler struct {
	service *expenseService
}

func newHandler(service *expenseService) *expenseHandler {
	return &expenseHandler{
		service: service,
	}
}

type expenseRequest struct {
	Name string `json:"name"`
	Value float32 `json:"value"`
	Category string `json:"category"`
}

func (eh *expenseHandler) CreateExpenseHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("token").(utils.Claims)
	if !ok {
		http.Error(w, "access failure", http.StatusNotFound)
		return
	}

	var expenseReq expenseRequest

	if err := json.NewDecoder(r.Body).Decode(&expenseReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	expense, err := eh.service.CreateExpenseService(r.Context(), expenseReq.Name, expenseReq.Category, user.UserId, expenseReq.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(expense)
}

func (eh *expenseHandler) GetExpenseByIdHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("token").(utils.Claims)
	if !ok {
		http.Error(w, "access failure", http.StatusNotFound)
		return
	}

	idStr := r.PathValue("id")
	expenseId, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	expense, err := eh.service.GetExpenseByIdService(r.Context(), expenseId, user.UserId)
	if err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expense)
}

func (eh *expenseHandler) GetAllExpensesHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("token").(utils.Claims)
	if !ok {
		http.Error(w, "access failure", http.StatusNotFound)
		return
	}

	expenses, err := eh.service.GetAllExpensesService(r.Context(), user.UserId)
	if err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expenses)
}

func (eh *expenseHandler) UpdateExpenseHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("token").(utils.Claims)
	if !ok {
		http.Error(w, "access failure", http.StatusNotFound)
		return
	}

	idStr := r.PathValue("id")
	expenseId, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var expenseReq expenseRequest
	
	if err := json.NewDecoder(r.Body).Decode(&expenseReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
		
	expense, err := eh.service.UpdateExpenseService(r.Context(), expenseId, user.UserId, expenseReq.Name, expenseReq.Category, expenseReq.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expense)
}

func (eh *expenseHandler) DeleteExpenseHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("token").(utils.Claims)
	if !ok {
		http.Error(w, "access failure", http.StatusNotFound)
		return
	}

	idStr := r.PathValue("id")
	expenseId, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	if err := eh.service.DeleteExpenseService(r.Context(), expenseId, user.UserId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("expense deleted")
}

func (eh *expenseHandler) FilterByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("token").(utils.Claims)
	if !ok {
		http.Error(w, "access failure", http.StatusNotFound)
		return
	}

	var expenseReq expenseRequest
	if err := json.NewDecoder(r.Body).Decode(&expenseReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	expenses, err := eh.service.FilterByCategoryService(r.Context(), expenseReq.Category, user.UserId)
	if err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expenses)
}

func (eh *expenseHandler) FilterByDateHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("token").(utils.Claims)
	if !ok {
		http.Error(w, "access failure", http.StatusNotFound)
		return
	}

	type dateReq struct {
		DateFrom string `json:"datefrom"`
		DateTo string `json:"dateto"`
	}

	var dates dateReq
	if err := json.NewDecoder(r.Body).Decode(&dates); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	expenses, err := eh.service.FilterByDateService(r.Context(), dates.DateFrom, dates.DateTo, user.UserId)
	if err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expenses)
}
