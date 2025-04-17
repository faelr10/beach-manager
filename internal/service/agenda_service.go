package service

import (
	"go/beach-manager/internal/domain"
	"go/beach-manager/internal/dto"
)

type AgendaService struct {
	repository domain.AgendaRepository
}

func NewAgendaService(agendaRepository domain.AgendaRepository) *AgendaService {
	return &AgendaService{repository: agendaRepository}
}

func (ag *AgendaService) CreateAgenda(input dto.CreateAgendaInput) (*dto.AgendaOutput, error) {
	agenda := dto.ToAgenda(input)

	// Verificar se há conflito de horário
	agendas, err := ag.repository.GetAllByUserID(agenda.UserID)
	if err != nil {
		return nil, err
	}

	for _, a := range agendas {
		if a.Date == agenda.Date {
			// Corrigir lógica de conflito
			if agenda.StartTime < a.EndTime && agenda.EndTime > a.StartTime {
				return nil, domain.ErrAgendaConflict
			}
		}
	}

	err = ag.repository.Create(agenda)
	if err != nil {
		return nil, err
	}

	output := dto.FromAgenda(agenda)
	return &output, nil
}


func (ag *AgendaService) GetAgendaByID(id string) (*dto.AgendaOutput, error) {
	agenda, err := ag.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	output := dto.FromAgenda(agenda)
	return &output, nil
}

func (ag *AgendaService) GetAllAgendas() ([]*dto.AgendaOutput, error) {
	agendas, err := ag.repository.GetAll()
	if err != nil {
		return nil, err
	}

	outputs := make([]*dto.AgendaOutput, len(agendas))
	for i, agenda := range agendas {
		output := dto.FromAgenda(agenda)
		outputs[i] = &output
	}

	return outputs, nil
}

func (ag *AgendaService) GetAllAgendasByUserID(userID string) ([]*dto.AgendaOutput, error) {
	agendas, err := ag.repository.GetAllByUserID(userID)
	if err != nil {
		return nil, err
	}
	outputs := make([]*dto.AgendaOutput, len(agendas))
	for i, agenda := range agendas {
		output := dto.FromAgenda(agenda)
		outputs[i] = &output
	}
	return outputs, nil
}

func (ag *AgendaService) UpdateAgenda(input dto.UpdateAgendaInput) (*dto.AgendaOutput, error) {

	agenda, err := ag.repository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	agenda.ClientName = input.ClientName
	agenda.Date = input.Date
	agenda.StartTime = input.StartTime
	agenda.EndTime = input.EndTime

	err = ag.repository.Update(agenda)
	if err != nil {
		return nil, err
	}

	output := dto.FromAgenda(agenda)
	return &output, nil

}

func (ag *AgendaService) DeleteAgenda(id string) error {
	err := ag.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
