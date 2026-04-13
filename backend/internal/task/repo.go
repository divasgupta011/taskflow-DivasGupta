package task

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, t *Task) error {
	query := `
		INSERT INTO tasks (title, description, status, priority, project_id, assignee_id, due_date)
		VALUES ($1,$2,$3,$4,$5,$6,$7)
		RETURNING id, created_at, updated_at
	`

	return r.db.QueryRowContext(ctx, query,
		t.Title, t.Description, t.Status, t.Priority,
		t.ProjectID, t.AssigneeID, t.DueDate,
	).Scan(&t.ID, &t.CreatedAt, &t.UpdatedAt)
}

func (r *Repository) GetByProject(ctx context.Context, projectID, status, assignee string) ([]Task, error) {
	query := `
		SELECT id, title, description, status, priority, project_id, assignee_id, due_date, created_at, updated_at
		FROM tasks
		WHERE project_id = $1
	`
	args := []interface{}{projectID}
	i := 2

	if status != "" {
		query += fmt.Sprintf(" AND status = $%d", i)
		args = append(args, status)
		i++
	}

	if assignee != "" {
		query += fmt.Sprintf(" AND assignee_id = $%d", i)
		args = append(args, assignee)
		i++
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		err := rows.Scan(
			&t.ID, &t.Title, &t.Description, &t.Status,
			&t.Priority, &t.ProjectID, &t.AssigneeID,
			&t.DueDate, &t.CreatedAt, &t.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (r *Repository) GetByID(ctx context.Context, id string) (*Task, error) {
	query := `
		SELECT id, title, description, status, priority, project_id, assignee_id, due_date, created_at, updated_at
		FROM tasks WHERE id = $1
	`

	var t Task
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&t.ID, &t.Title, &t.Description, &t.Status,
		&t.Priority, &t.ProjectID, &t.AssigneeID,
		&t.DueDate, &t.CreatedAt, &t.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (r *Repository) Update(ctx context.Context, t *Task) error {
	query := `
		UPDATE tasks
		SET title=$1, description=$2, status=$3, priority=$4, assignee_id=$5, due_date=$6, updated_at=NOW()
		WHERE id=$7
	`

	_, err := r.db.ExecContext(ctx, query,
		t.Title, t.Description, t.Status, t.Priority,
		t.AssigneeID, t.DueDate, t.ID,
	)
	return err
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM tasks WHERE id=$1", id)
	return err
}
