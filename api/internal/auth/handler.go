package auth

import (
	"encoding/json"
	"expense-api/api/internal/utils"
	"net/http"
	"time"
)

type authHandler struct {
	service *authService	
}

func newHandler(service *authService) *authHandler {
	return &authHandler{
		service: service,
	}
}

type authRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func (ah *authHandler) LogInHandler(w http.ResponseWriter, r *http.Request) {
	var userReq authRequest
	
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	user, err := ah.service.LogInService(r.Context(), userReq.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !utils.CheckPassword(userReq.Password, user.UserPassword) {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateToken(user.UserId, user.UserName, user.UserEmail)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: token,
		Path: "/",
		Expires: time.Now().Add(24 * time.Hour),
	})

	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Authorization", token)
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
		"message": "User successfully logged in",
	})
}

func (ah *authHandler) LogOutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		MaxAge: -1,
		Path: "/",
		Expires: time.Now().Add(-100 * time.Hour),
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("User successfully logged out")
}
