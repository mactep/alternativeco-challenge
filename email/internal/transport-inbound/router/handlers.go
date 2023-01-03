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

type EmailRequest struct {
	Email string `json:"email"`
}

func (h emailHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var emailRequest EmailRequest
	err := json.NewDecoder(r.Body).Decode(&emailRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if emailRequest.Email == "" {
		http.Error(w, "no email provided", http.StatusBadRequest)
		return
	}

	if !features.IsValidEmail(emailRequest.Email) {
		http.Error(w, "invalid email provided", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	emailRegistry, err := h.service.Create(ctx, emailRequest.Email)

	if err != nil {
		http.Error(w, fmt.Sprintf("error creating email: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

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
