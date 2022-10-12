package tickets

import (
	"desafio-goweb-saraisanjuan/internal/domain"
)

type Service interface {
	GetTotalTickets(destination string) ([]domain.Ticket, error)
	AverageDestination(destination string) (float64, error)
}
type service struct {
	repository Repository
}

func NewServices(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetTotalTickets(destination string) ([]domain.Ticket, error) {
	return s.repository.GetTicketByDestination(destination)
}

func (s *service) AverageDestination(destination string) (float64, error) {
	listtotal, err := s.repository.GetAll()
	if err != nil {
		return 0, err
	}
	listdestination, err := s.repository.GetTicketByDestination(destination)
	if err != nil {
		return 0, err
	}
	return float64((len(listdestination) * 100)) / float64(len(listtotal)), nil
}
