package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Polidoro-root/go-expert/7_apis/internal/dto"
	"github.com/Polidoro-root/go-expert/7_apis/internal/entity"
	"github.com/Polidoro-root/go-expert/7_apis/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{UserDB: userDB}
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("JWTExpiresIn").(int)

	var user dto.GetJWTInput

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.UserDB.FindByEmail(user.Email)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix,
	})

	acessToken := struct {
		AcessToken string `json:"access_token"`
	}{AcessToken: tokenString}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(acessToken)

}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.UserDB.Create(u)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
