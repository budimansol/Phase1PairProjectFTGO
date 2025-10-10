package service

import (
	"time"
	"github.com/budimansol/pairproject/internal/model"
	"github.com/budimansol/pairproject/internal/repository"
)

type TransactionService struct {
	repo *repository.TransactionRepository
}

func NewTransactionService(repo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) CreateTransaction(tx *model.Transaction) (int, error) {
	// Hitung total
	var total float64
	for i := range tx.Items {
		tx.Items[i].Subtotal = float64(tx.Items[i].Quantity) * tx.Items[i].Price
		total += tx.Items[i].Subtotal
	}
	tx.TotalAmount = total
	tx.TransactionDate = time.Now()

	id, err := s.repo.Create(tx)
	if err != nil {
		return 0, err
	}

	for i := range tx.Items {
		tx.Items[i].TransactionID = id
		if err := s.repo.AddItem(&tx.Items[i]); err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (s *TransactionService) GetAllTransactions() ([]model.Transaction, error) {
	return s.repo.GetAll()
}

func (s *TransactionService) DeleteTransaction(id int) error {
	return s.repo.Delete(id)
}
