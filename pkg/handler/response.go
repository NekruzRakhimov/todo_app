package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

type dataResponse struct {
	Data map[string]interface{} `json:"data"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

func newErrResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(errorResponse{Message: message}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func newStatusResponse(w http.ResponseWriter, status string) {
	body, err := json.Marshal(statusResponse{Status: status})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	_, err = w.Write(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func newDataResponse(w http.ResponseWriter, data dataResponse) {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
