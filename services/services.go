package services

import (
	"github.com/dbsimmons64/go-beans/models"
	"github.com/dbsimmons64/go-beans/repos"
)

type TransactionService interface {
	ListTransactions() ([]models.Transaction, error)
}

type TransactionServiceImpl struct {
	Repo repos.TransactionRepository
}

func (t TransactionServiceImpl) ListTransactions() ([]models.Transaction, error) {
	return t.Repo.All()
}
