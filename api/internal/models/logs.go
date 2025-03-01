package models

import (
	"context"
	"time"

	"github.com/Vladisluv228/ApiBdGoExample/api/internal/db"
)

type Log struct {
	ID int `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	EditedAt     time.Time `json:"edited_at"`
	Message string `json:"message"`
}

type LogInput struct {
	ID int `json:"id"`
	Message string
}

func GetLogs(ctx context.Context) ([]Log, error) {
	query := `
		SELECT id, created_at, edited_at, message
		FROM logs
	`
	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var logs []Log
	for rows.Next() {
		var log Log
		err := rows.Scan(&log.ID, &log.CreatedAt, &log.EditedAt, &log.Message)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return logs, nil
}

func CreateLog(ctx context.Context, input LogInput) (*Log, error) {
	var log Log
	query := `
		INSERT INTO logs (id, created_at, edited_at, message)
		VALUES ($1, $2, $3, $4)
		RETURNING id, message
	`

	err := db.QueryRow(
		ctx,
		query,
		input.ID,
		time.Now(),
		time.Now(),
		input.Message,
	).Scan(&log.ID, log.Message)

	if err != nil {
		return nil, err
	}

	log.CreatedAt = time.Now()
	log.EditedAt = time.Now()
	return &log, err
}