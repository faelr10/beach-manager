package domain

import (
	"time"

	"github.com/google/uuid"
)

type Agenda struct {
	ID         string
	UserID     string
	ClientName string
	Date       string
	StartTime  string
	EndTime    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewAgenda(userID, client_name, date, startTime, endTime string) *Agenda {
	return &Agenda{
		ID:         uuid.New().String(),
		UserID:     userID,
		ClientName: client_name,
		Date:       date,
		StartTime:  startTime,
		EndTime:    endTime,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
