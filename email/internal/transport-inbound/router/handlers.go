package router

import (
	"fmt"
	"net/http"

	"github.com/gothunder/thunder/pkg/router"
)

func NewEmailHandler() router.HandlerOutput {
	return router.HandlerOutput{
		Handler: emailHandler{},
	}
}

type emailHandler struct{}

func (h emailHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello email")
}

func (h emailHandler) Method() string {
	return "Get"
}

func (h emailHandler) Pattern() string {
	return "/email"
}
