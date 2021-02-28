package controller

import (
	"log"
	"net/http"

	todoerr "github.com/manabie-com/togo/internal/httperror"
	"github.com/manabie-com/togo/internal/services"
)

// Handler of server
type Handler struct {
	Identity *services.IdentityService
	ToDo     *services.ToDoService
}

func (h *Handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "*")
	resp.Header().Set("Content-Type", "application/json")

	if req.Method == http.MethodOptions {
		resp.WriteHeader(http.StatusOK)
		return
	}

	switch req.URL.Path {
	case "/login":
		h.Identity.GetAuthToken(resp, req)
		return

	case "/tasks":
		var ok bool
		req, ok = h.Identity.ValidToken(req)
		if !ok {
			todoerr.WrapError(resp, http.StatusUnauthorized, "Unauthorized")
			return
		}

		switch req.Method {
		case http.MethodGet:
			h.ToDo.ListTasks(resp, req)
		case http.MethodPost:
			h.ToDo.AddTask(resp, req)
		case http.MethodDelete:
			h.ToDo.DeleteTask(resp, req)
		}
		return
	default:
		todoerr.WrapError(resp, http.StatusNotFound, "Not fould")
		return
	}
}
