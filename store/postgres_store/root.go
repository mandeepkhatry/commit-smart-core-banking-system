package postgres_store

import (
	"database/sql"

	"github.com/commit-smart-core-banking-system/config"
	_ "github.com/lib/pq"
)

//Adapter
type PostGres struct {
	db *sql.DB
}

func NewClient(dbSource string) (*PostGres, error) {
	conn, err := sql.Open("postgres", config.AppConfiguration.DbSource)
	if err != nil {
		return &PostGres{}, err
	}

	if err := conn.Ping(); err != nil {
		return &PostGres{}, err
	}

	return &PostGres{db: conn}, nil

}

func (p *PostGres) CloseClient() error {
	return p.db.Close()
}
