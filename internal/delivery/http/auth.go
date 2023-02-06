package http

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"practice/internal/domain"
)

func (h *Handler) signUp(c gin.Context) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logError("signUp", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var input domain.SignUpInput
	if err = json.Unmarshal(reqBytes, &input); err != nil {
		logError("signUp", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := input.Validate(); err != nil {
		logError("signUp", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.user

	w.WriteHeader(http.StatusOK)
}
