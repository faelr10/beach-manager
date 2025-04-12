package domain

import (
	"time"

	"github.com/google/uuid"
)

type Agenda struct {
	ID        string
	UserID    string
	Date      string
	StartTime string
	EndTime   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAgenda(userID, date, startTime, endTime string) *Agenda {
	return &Agenda{
		ID:        uuid.New().String(),
		UserID:    userID,
		Date:      date,
		StartTime: startTime,
		EndTime:   endTime,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
