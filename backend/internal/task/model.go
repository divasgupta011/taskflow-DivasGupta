package task

import "time"

type Task struct {
	ID          string
	Title       string
	Description *string
	Status      string
	Priority    string
	ProjectID   string
	AssigneeID  *string
	DueDate     *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
