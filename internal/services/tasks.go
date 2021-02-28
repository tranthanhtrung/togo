package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

	todoerr "github.com/manabie-com/togo/internal/httperror"
	"github.com/manabie-com/togo/internal/storages"
	todosql "github.com/manabie-com/togo/internal/storages/sql"
	"github.com/manabie-com/togo/utils"
)

// ToDoService implement HTTP server
type ToDoService struct {
	MaxTasks int
	Store    *todosql.Database
}

// NewToDoService create a ToDoService
func NewToDoService(maxTasks int, store *todosql.Database) *ToDoService {
	return &ToDoService{
		MaxTasks: maxTasks,
		Store:    store,
	}
}

// ListTasks get list task of a user
func (s *ToDoService) ListTasks(resp http.ResponseWriter, req *http.Request) {
	id, _ := utils.UserIDFromCtx(req.Context())
	tasks, err := s.Store.RetrieveTasks(
		req.Context(),
		sql.NullString{
			String: id,
			Valid:  true,
		},
		utils.ValueFromHTTPRequest(req, "created_date"),
	)

	if err != nil {
		todoerr.WrapError(resp, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(resp).Encode(map[string][]*storages.Task{
		"data": tasks,
	})
}

// AddTask insert a task of a user
func (s *ToDoService) AddTask(resp http.ResponseWriter, req *http.Request) {
	task := &storages.Task{}
	err := json.NewDecoder(req.Body).Decode(task)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	now := time.Now()
	userID, _ := utils.UserIDFromCtx(req.Context())
	task = storages.NewTask(
		uuid.New().String(),
		task.Content,
		userID,
		now.Format("2006-01-02"),
	)

	tasks, err := s.Store.RetrieveTasks(
		req.Context(),
		sql.NullString{
			String: task.UserID,
			Valid:  true,
		},
		sql.NullString{
			String: task.CreatedDate,
			Valid:  true,
		},
	)
	if err != nil {
		todoerr.WrapError(resp, http.StatusInternalServerError, err.Error())
		return
	}
	if len(tasks) >= s.MaxTasks {
		todoerr.WrapError(resp, http.StatusBadRequest, fmt.Sprintf("Limit %d task per day", s.MaxTasks))
		return
	}

	err = s.Store.AddTask(req.Context(), task)
	if err != nil {
		todoerr.WrapError(resp, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(resp).Encode(map[string]*storages.Task{
		"data": task,
	})
}

// DeleteTask delete a task of a user
func (s *ToDoService) DeleteTask(resp http.ResponseWriter, req *http.Request) {
	task := &storages.Task{}
	err := json.NewDecoder(req.Body).Decode(task)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = s.Store.DeleteTask(req.Context(), task.ID)
	if err != nil {
		todoerr.WrapError(resp, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(resp).Encode(map[string]string{
		"id": task.ID,
	})
}
