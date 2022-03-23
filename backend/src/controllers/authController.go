package controllers

import (
	"encoding/json"
	"github.com/qndaa/backend/src/services"
	"net/http"
)

type AuthController struct {
	userService *services.UserService
}

var AuthInstanceController AuthController

func init() {
	AuthInstanceController = AuthController{
		userService: &services.UserInstanceService,
	}
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	c.userService.Login(w, data["username"], data["password"])
}
