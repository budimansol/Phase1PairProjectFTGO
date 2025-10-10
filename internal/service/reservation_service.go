package service

import (
	"github.com/budimansol/pairproject/internal/model"
	"github.com/budimansol/pairproject/internal/repository"
)

type ReservationService struct {
	repo *repository.ReservationRepository
}

func NewReservationService(repo *repository.ReservationRepository) *ReservationService {
	return &ReservationService{repo: repo}
}

func (s *ReservationService) CreateReservation(res model.Reservation) error {
	return s.repo.Create(res)
}

func (s *ReservationService) GetAllReservations() ([]model.Reservation, error) {
	return s.repo.GetAll()
}
