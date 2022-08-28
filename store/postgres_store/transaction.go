package postgres_store

import (
	"fmt"

	"github.com/commit-smart-core-banking-system/logger"
	"github.com/commit-smart-core-banking-system/model"
)

const createNonTransferTxn = `
INSERT INTO txn(
	account_id,
	transaction_type,
	amount_type,
	amount,
	description,
	current_balance
) VALUES (
	$1, $2, $3, $4, $5, $6
) RETURNING id, account_id, transaction_type, amount_type, amount, description, current_balance, created_at
`

const createTransferTxn = `
INSERT INTO txn(
	account_id,
	transaction_type,
	amount_type,
	amount,
	description,
	corresponding_account_id,
	current_balance
) VALUES (
	$1, $2, $3, $4, $5, $6, $7
) RETURNING id, account_id, transaction_type, amount_type, amount, description, corresponding_account_id, current_balance, created_at
`

func (p *PostGres) CreateTxn(txn model.TxnRequest) (model.TxnResponse, error) {
	var txnResponse model.TxnResponse
	if txn.TransactionType == model.TransactionType.Transfer {
		row := p.db.QueryRow(createTransferTxn, txn.AccountId, txn.TransactionType, txn.AmountType, txn.Amount, txn.Description, txn.CorrespondingAccountId, txn.CurrentBalance)
		err := row.Scan(
			&txnResponse.Id,
			&txnResponse.AccountId,
			&txnResponse.TransactionType,
			&txnResponse.AmountType,
			&txnResponse.Amount,
			&txnResponse.Description,
			&txnResponse.CorrespondingAccountId,
			&txnResponse.CurrentBalance,
			&txnResponse.CreatedAt,
		)
		return txnResponse, err
	}

	row := p.db.QueryRow(createNonTransferTxn, txn.AccountId, txn.TransactionType, txn.AmountType, txn.Amount, txn.Description, txn.CurrentBalance)
	err := row.Scan(
		&txnResponse.Id,
		&txnResponse.AccountId,
		&txnResponse.TransactionType,
		&txnResponse.AmountType,
		&txnResponse.Amount,
		&txnResponse.Description,
		&txnResponse.CurrentBalance,
		&txnResponse.CreatedAt,
	)
	return txnResponse, err

}

func (p *PostGres) ListTxn(filter model.TransactionFilter) (model.ListTxnResponse, error) {

	var listTxnQuery = "SELECT id,account_id,transaction_type,amount_type,amount,description,current_balance,created_at FROM txn"
	listTxnQuery = GenerateListTransactionQuery(filter, listTxnQuery)

	var list model.ListTxnResponse

	rows, err := p.db.Query(listTxnQuery)
	if err != nil {
		return model.ListTxnResponse{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var txn model.TxnResponse
		if err := rows.Scan(
			&txn.Id,
			&txn.AccountId,
			&txn.TransactionType,
			&txn.AmountType,
			&txn.Amount,
			&txn.Description,
			&txn.CurrentBalance,
			&txn.CreatedAt,
		); err != nil {
			return model.ListTxnResponse{}, err
		}
		list.Txn = append(list.Txn, txn)
	}

	if err := rows.Close(); err != nil {
		return model.ListTxnResponse{}, err
	}
	if err := rows.Err(); err != nil {
		return model.ListTxnResponse{}, err
	}
	return list, nil

}

func GenerateListTransactionQuery(filter model.TransactionFilter, query string) string {
	query += fmt.Sprintf(" WHERE account_id=%d", filter.AccountId)
	if filter.Type != model.FilterParam.All {
		query += fmt.Sprintf(" AND transaction_type='%s'", filter.Type)
	}
	logger.Debug("Txn Query : ", logger.Field("query", query))

	query += fmt.Sprintf(" ORDER BY created_at %s LIMIT %d OFFSET %d", filter.Order, filter.Limit, filter.Offset)
	return query
}
