package storages

// Task reflects tasks in DB
type Task struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	UserID      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}

// NewTask create a Task
func NewTask(id, content, userID, createdDate string) *Task {
	return &Task{
		ID:          id,
		Content:     content,
		UserID:      userID,
		CreatedDate: createdDate,
	}
}

// User reflects users data from DB
type User struct {
	ID       string
	Password string
}
