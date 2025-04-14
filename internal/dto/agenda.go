package dto

import "go/beach-manager/internal/domain"

type CreateAgendaInput struct {
	UserID     string `json:"user_id"`
	ClientName string `json:"client_name"`
	Date       string `json:"date"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
}

type AgendaOutput struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	ClientName string `json:"client_name" validate:"required"`
	Date       string `json:"date"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

func ToAgenda(input CreateAgendaInput) *domain.Agenda {
	return domain.NewAgenda(input.UserID, input.ClientName, input.Date, input.StartTime, input.EndTime)
}

func FromAgenda(agenda *domain.Agenda) AgendaOutput {
	return AgendaOutput{
		ID:         agenda.ID,
		UserID:     agenda.UserID,
		ClientName: agenda.ClientName,
		Date:       agenda.Date,
		StartTime:  agenda.StartTime,
		EndTime:    agenda.EndTime,
		CreatedAt:  agenda.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:  agenda.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}
