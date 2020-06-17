package service

import (

	"database/sql"

	_ "github.com/lib/pq"

	model "../model"
)

type IDomainService interface {
	FetchDomain(db *sql.DB, domain string) (*model.Domain, error)
}

func FetchDomain(db *sql.DB, domain string) (*model.Domain, error) { //Here formatted to just return domain names, not times
	return &model.Domain{}, nil
}

