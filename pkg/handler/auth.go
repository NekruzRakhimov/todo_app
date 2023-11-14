package handler

import (
	"encoding/json"
	"github.com/NekruzRakhimov/todo_app/models"
	"github.com/NekruzRakhimov/todo_app/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"io"
	"net/http"
)

type Auth struct {
	services *service.Service
	nats     *nats.Conn
}

func NewAuth(services *service.Service, nats *nats.Conn) *Auth {
	return &Auth{services: services, nats: nats}
}

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body models.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (a *Auth) signUp(w http.ResponseWriter, r *http.Request) {
	var input models.User
	body, err := io.ReadAll(r.Body)
	if err != nil {
		newErrResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = json.Unmarshal(body, &input); err != nil {
		newErrResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := a.services.Authorization.CreateUser(input)
	if err != nil {
		newErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newDataResponse(w, dataResponse{Data: map[string]interface{}{"id": id}})
}

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body models.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (a *Auth) signIn(w http.ResponseWriter, r *http.Request) {
	var input models.SignInInput
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(body, &input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := a.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newDataResponse(w, dataResponse{Data: map[string]interface{}{"token": token}})
}

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body models.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body models.SignInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input models.SignInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
