package postgres_store

import (
	"github.com/commit-smart-core-banking-system/model"
)

const createCustomer = `
INSERT INTO customer(
	name,
	address,
	phone_number,
	email_address,
	password_hash
) VALUES (
	$1, $2, $3, $4, $5
) RETURNING id, name, address, phone_number, email_address, created_at
`

func (p *PostGres) CreateCustomer(customer model.Customer) (model.CustomerResponse, error) {
	var customerResponse model.CustomerResponse
	row := p.db.QueryRow(createCustomer, customer.Name, customer.Address, customer.PhoneNumber, customer.EmailAddress, customer.PasswordHash)
	err := row.Scan(
		&customerResponse.Id,
		&customerResponse.Name,
		&customerResponse.Address,
		&customerResponse.PhoneNumber,
		&customerResponse.EmailAddress,
		&customerResponse.CreatedAt,
	)
	return customerResponse, err

}

const fetchCustomer = `
SELECT * FROM customer
WHERE email_address = $1 LIMIT 1
`

func (p *PostGres) FetchCustomer(email string) (model.CustomerData, error) {
	row := p.db.QueryRow(fetchCustomer, email)
	var customer model.CustomerData
	err := row.Scan(
		&customer.Id,
		&customer.Name,
		&customer.Address,
		&customer.PhoneNumber,
		&customer.EmailAddress,
		&customer.PasswordHash,
		&customer.CreatedAt,
	)
	return customer, err
}
