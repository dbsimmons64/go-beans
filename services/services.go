package services

import (
	"net/url"

	"github.com/dbsimmons64/go-beans/models"
	"github.com/dbsimmons64/go-beans/repos"
)

type TransactionService interface {
	ListTransactions() ([]models.Transaction, error)
	CreateTransaction(data url.Values) error
}

type TransactionServiceImpl struct {
	Repo repos.TransactionRepository
}

func (t TransactionServiceImpl) ListTransactions() ([]models.Transaction, error) {
	return t.Repo.All()
}

func (t TransactionServiceImpl) CreateTransaction(data url.Values) error {
	return t.Repo.Insert(data)
}
