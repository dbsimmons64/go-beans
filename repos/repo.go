package repos

import (
	"database/sql"

	"github.com/dbsimmons64/go-beans/models"
)

type TransactionRepository interface {
	All() ([]models.Transaction, error)
}

type TransactionRepositoryDB struct {
	DB *sql.DB
}

func (t TransactionRepositoryDB) All() ([]models.Transaction, error) {

	stmt := `SELECT id, txn_date, who, description, payee, amount, category FROM transactions ORDER BY id DESC`
	rows, err := t.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	transactions := []models.Transaction{}

	for rows.Next() {
		p := models.Transaction{}
		err := rows.Scan(&p.Id, &p.Txn_date, &p.Who, &p.Description, &p.Payee, &p.Amount, &p.Category)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, p)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return transactions, nil

}
