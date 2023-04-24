package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ReygaFitra/inc-final-project.git/model"
)

var now = time.Now().UTC().Truncate(time.Minute)
var waktu = now.Format("2006-01-02 15:04")

type TransactionRepository interface {
	CreateDepositBank(tx *model.TransactionBank) error
	CreateDepositCard(tx *model.TransactionCard) error
	CreateWithdrawal(tx *model.TransactionWithdraw) error
	CreateTransfer(tx *model.TransactionTransfer) (any, error)
	CreateRedeem(tx *model.TransactionPoint) error
	GetAllPoint() ([]*model.PointExchange, error)
}

type transactionRepository struct {
	db *sql.DB
}

func (r *transactionRepository) CreateDepositBank(tx *model.TransactionBank) error {
	query := "INSERT INTO tx_transaction (transaction_type, sender_id, bank_account_id, amount, transaction_date) VALUES ($1, $2, $3, $4, $5)"
	_, err := r.db.Exec(query, "Deposit Bank", tx.SenderID, tx.BankAccountID, tx.Amount, waktu)
	if err != nil {
		return err
	}

	return nil
}

func (r *transactionRepository) CreateDepositCard(tx *model.TransactionCard) error {
	query := "INSERT INTO tx_transaction (transaction_type, sender_id, card_id, amount, transaction_date) VALUES ($1, $2, $3, $4, $5)"
	_, err := r.db.Exec(query, "Deposit Card", tx.SenderID, tx.CardID, tx.Amount, waktu)
	if err != nil {
		return err
	}

	return nil
}

func (r *transactionRepository) CreateWithdrawal(tx *model.TransactionWithdraw) error {
	query := "INSERT INTO tx_transaction (transaction_type, bank_account_id, sender_id, amount, transaction_date) VALUES ($1, $2, $3, $4,$5)"
	_, err := r.db.Exec(query, "Withdraw", tx.BankAccountID, tx.SenderID, tx.Amount, waktu)
	if err != nil {
		return err
	}

	return nil
}

func (r *transactionRepository) CreateTransfer(tx *model.TransactionTransfer) (any, error) {
	query := "INSERT INTO tx_transaction (transaction_type, sender_id, recipient_id, amount, transaction_date) VALUES ($1, $2, $3, $4, $5)"
	_, err := r.db.Exec(query, "Transfer", tx.SenderID, tx.RecipientID, tx.Amount, waktu)
	if err != nil {
		return nil, fmt.Errorf("failed to create data: %v", err)
	}

	return tx, nil
}

func (r *transactionRepository) CreateRedeem(tx *model.TransactionPoint) error {
	query := "INSERT INTO tx_transaction (transaction_type, sender_id, pe_id, point, transaction_date) VALUES ($1, $2, $3, $4, $5)"
	_, err := r.db.Exec(query, "Redeem", tx.SenderID, tx.PointExchangeID, tx.Point, waktu)
	if err != nil {
		return err
	}

	return nil
}

// Get all point exchanges
func (r *transactionRepository) GetAllPoint() ([]*model.PointExchange, error) {
	query := "SELECT pe_id, reward, price FROM mst_point_exchange"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pointExchanges []*model.PointExchange
	for rows.Next() {
		pe := &model.PointExchange{}
		err := rows.Scan(&pe.PE_ID, &pe.Reward, &pe.Price)
		if err != nil {
			return nil, err
		}
		pointExchanges = append(pointExchanges, pe)
	}

	return pointExchanges, nil
}

func NewTxRepository(db *sql.DB) TransactionRepository {
	repo := new(transactionRepository)
	repo.db = db
	return repo
}
