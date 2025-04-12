package dto

import "go/beach-manager/internal/domain"

type CreateAgendaInput struct {
	UserID    string `json:"user_id"`
	Date      string `json:"date"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type AgendaOutput struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Date      string `json:"date"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ToAgenda(input CreateAgendaInput) *domain.Agenda {
	return domain.NewAgenda(input.UserID, input.Date, input.StartTime, input.EndTime)
}

func FromAgenda(agenda *domain.Agenda) AgendaOutput {
	return AgendaOutput{
		ID:        agenda.ID,
		UserID:    agenda.UserID,
		Date:      agenda.Date,
		StartTime: agenda.StartTime,
		EndTime:   agenda.EndTime,
		CreatedAt: agenda.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: agenda.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}
