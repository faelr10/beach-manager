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

	err := ag.repository.Create(agenda)
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