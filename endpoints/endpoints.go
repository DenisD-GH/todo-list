package endpoints

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func GetTasks(conn *pgx.Conn, ctx context.Context) ([]Task, error) {
	rows, err := conn.Query(ctx, "SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func CreateTask(conn *pgx.Conn, ctx context.Context, task Task) error {
	_, err := conn.Exec(ctx, `
		INSERT INTO tasks (title, description, status, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())`,
		task.Title, task.Description, task.Status)
	return err
}

func UpdateTask(conn *pgx.Conn, ctx context.Context, task Task) error {
	_, err := conn.Exec(ctx, `
        UPDATE tasks
        SET title = $1, description = $2, status = $3, updated_at = NOW()
        WHERE id = $4`,
		task.Title, task.Description, task.Status, task.ID)
	return err
}

func DeleteTask(conn *pgx.Conn, ctx context.Context, id int) error {
	_, err := conn.Exec(ctx, `DELETE FROM tasks WHERE id = $1`, id)
	return err
}
