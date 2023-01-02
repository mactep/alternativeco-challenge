package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gothunder/thunder/pkg/router"
	"github.com/mactep/alternativeco-challenge/email/internal/features"
)

func NewEmailHandler(emailService features.EmailService) router.HandlerOutput {
	return router.HandlerOutput{
		Handler: emailHandler{emailService},
	}
}

type emailHandler struct {
	service features.EmailService
}

func (h emailHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	email, ok := req["email"].(string)
	if !ok {
		http.Error(w, "email is not a string", http.StatusBadRequest)
		return
	}

	if email == "" {
		http.Error(w, "no email provided", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	emailRegistry, err := h.service.Create(ctx, email)

	if err != nil {
		http.Error(w, fmt.Sprintf("error creating email: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// marshal emailRegistry to JSON and write to response
	jsonEmailRegistry, err := json.Marshal(emailRegistry)
	if err != nil {
		http.Error(w, fmt.Sprintf("error marshalling emailRegistry: %s", err), http.StatusInternalServerError)
		return
	}

	w.Write(jsonEmailRegistry)
}

func (h emailHandler) Method() string {
	return http.MethodPost
}

func (h emailHandler) Pattern() string {
	return "/email"
}
