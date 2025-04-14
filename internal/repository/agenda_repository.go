package repository

import (
	"database/sql"
	"go/beach-manager/internal/domain"
)

type AgendaRepository struct {
	db *sql.DB
}

func NewAgendaRepository(db *sql.DB) *AgendaRepository {
	return &AgendaRepository{db: db}
}

func (r *AgendaRepository) Create(agenda *domain.Agenda) error {
	query := "INSERT INTO agendas (id, user_id, client_name, date, start_time, end_time) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := r.db.Exec(query, agenda.ID, agenda.UserID, agenda.ClientName, agenda.Date, agenda.StartTime, agenda.EndTime)
	if err != nil {
		return err
	}
	return nil
}

func (r *AgendaRepository) GetByID(id string) (*domain.Agenda, error) {
	query := "SELECT id, user_id, client_name, date, start_time, end_time FROM agendas WHERE id = $1"
	row := r.db.QueryRow(query, id)

	var agenda domain.Agenda
	err := row.Scan(&agenda.ID, &agenda.UserID, &agenda.ClientName, &agenda.Date, &agenda.StartTime, &agenda.EndTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrAgendaNotFound
		}
		return nil, err
	}

	return &agenda, nil
}

func (r *AgendaRepository) GetAll() ([]*domain.Agenda, error) {
	query := "SELECT id, user_id, client_name, date, start_time, end_time FROM agendas"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var agendas []*domain.Agenda
	for rows.Next() {
		var agenda domain.Agenda
		err := rows.Scan(&agenda.ID, &agenda.UserID, &agenda.ClientName, &agenda.Date, &agenda.StartTime, &agenda.EndTime)
		if err != nil {
			return nil, err
		}
		agendas = append(agendas, &agenda)
	}

	return agendas, nil
}
