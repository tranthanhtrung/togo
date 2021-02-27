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
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	"github.com/manabie-com/togo/utils"
)

// ToDoService implement HTTP server
type ToDoService struct {
	MaxTaskes int
	Store     *sqllite.LiteDB
}

// NewToDoService create a ToDoService
func NewToDoService(maxTaskes int, store *sqllite.LiteDB) *ToDoService {
	return &ToDoService{
		MaxTaskes: maxTaskes,
		Store:     store,
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
	if len(tasks) > s.MaxTaskes {
		todoerr.WrapError(resp, http.StatusBadRequest, fmt.Sprintf("Limit %d task per day", s.MaxTaskes))
		return
	}

	err = s.Store.AddTask(req.Context(), task)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]*storages.Task{
		"data": task,
	})
}
