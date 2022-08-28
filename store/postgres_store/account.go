package postgres_store

import (
	"github.com/commit-smart-core-banking-system/model"
)

const createAccount = `
INSERT INTO account(
	type,
	customer_id,
	balance
) VALUES (
	$1, $2, $3
) RETURNING id, type, customer_id, created_at
`

func (p *PostGres) CreateAccount(account model.Account) (model.AccountResponse, error) {
	var accountResponse model.AccountResponse
	row := p.db.QueryRow(createAccount, account.Type, account.CustomerId, account.Balance)
	err := row.Scan(
		&accountResponse.Id,
		&accountResponse.Type,
		&accountResponse.CustomerId,
		&accountResponse.CreatedAt,
	)
	return accountResponse, err

}

const listAccount = `
SELECT 
	a.id, 
	a.type, 
	c.name ,
	a.created_at
FROM account a INNER JOIN customer c 
ON a.customer_id=c.id 
WHERE type=$1  
ORDER BY c.name 
LIMIT $2 OFFSET $3
`

func (p *PostGres) ListAccount(filter model.AccountFilter) (model.ListAccountResponse, error) {

	var list model.ListAccountResponse
	rows, err := p.db.Query(listAccount, filter.Type, filter.Limit, filter.Offset)
	if err != nil {
		return model.ListAccountResponse{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var account model.AccountCustomer
		if err := rows.Scan(
			&account.Id,
			&account.Type,
			&account.Name,
			&account.CreatedAt,
		); err != nil {
			return model.ListAccountResponse{}, err
		}
		list.Accounts = append(list.Accounts, account)
	}

	if err := rows.Close(); err != nil {
		return model.ListAccountResponse{}, err
	}
	if err := rows.Err(); err != nil {
		return model.ListAccountResponse{}, err
	}
	return list, nil

}

const fetchAccount = `
SELECT id, type, balance, created_at FROM account
WHERE id = $1 LIMIT 1
`

func (p *PostGres) FetchAccount(id int64) (model.AccountResponse, error) {
	row := p.db.QueryRow(fetchAccount, id)
	var account model.AccountResponse
	err := row.Scan(
		&account.Id,
		&account.Type,
		&account.Balance,
		&account.CreatedAt,
	)
	return account, err
}

const updateAccount = `
UPDATE account 
SET balance=$1
WHERE id=$2
RETURNING id, type, customer_id, balance,created_at
`

func (p *PostGres) UpdateAccount(account model.AccountResponse) (model.AccountResponse, error) {
	var accountResponse model.AccountResponse
	row := p.db.QueryRow(updateAccount, account.Balance, account.Id)
	err := row.Scan(
		&accountResponse.Id,
		&accountResponse.Type,
		&accountResponse.CustomerId,
		&accountResponse.Balance,
		&accountResponse.CreatedAt,
	)
	return accountResponse, err
}

const fetchAccountByCustomerId = `
SELECT id, type, balance, created_at FROM account
WHERE customer_id = $1 LIMIT 1
`

func (p *PostGres) FetchAccountByCustomerId(id int64) (model.AccountResponse, error) {
	row := p.db.QueryRow(fetchAccountByCustomerId, id)
	var account model.AccountResponse
	err := row.Scan(
		&account.Id,
		&account.Type,
		&account.Balance,
		&account.CreatedAt,
	)
	return account, err
}
