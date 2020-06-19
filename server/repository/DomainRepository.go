package database

import (
	"database/sql"
	"log"
	database "../database"
)

var DOMAIN_TABLE string = "domain"

type IDomainRepository interface {
	CreateDomain(d *database.DomainDB) error
	FetchDomain(domain string) (database.DomainDB, error)
	UpdateDomain(d database.DomainDB) error
}

type domainRepository struct {
	db *sql.DB
}

func NewDomainRepository(db *sql.DB) domainRepository {
	return domainRepository{db: db}
}

func (r domainRepository) CreateDomain(d *database.DomainDB) error {
	sqlStm := `INSERT INTO `+DOMAIN_TABLE+` (name, sslgrade, previoussslgrade, searchedat) 
	VALUES ($1, $2, $3, NOW())`
	_, err := r.db.Exec(sqlStm, d.Name, d.SslGrade, d.PreviousSslGrade, d.SearchedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r domainRepository) FetchDomain(domain string) ([]database.DomainDB, error) {
	sqlStm := `SELECT * FROM `+DOMAIN_TABLE+` WHERE name='`+domain+`'`
	rows, err := r.db.Query(sqlStm)
	if err != nil {
		return *&[]database.DomainDB{}, err
	}
	defer rows.Close()
	var domains []database.DomainDB

	for rows.Next() {
		var d database.DomainDB
		if err := rows.Scan(&d.Name, &d.SslGrade, &d.PreviousSslGrade, &d.SearchedAt); err != nil {
			log.Println(err)
			continue
		}
		domains = append(domains, d)
	}
	return domains, nil
}

func (r domainRepository) UpdateDomain(d database.DomainDB) error {
	sqlStm := `UPDATE `+DOMAIN_TABLE+` SET SslGrade = ? , PreviousSslGrade = ? , SearchedAt = NOW() WHERE Name = ?`
	_, err := r.db.Exec(sqlStm,d.SslGrade,d.PreviousSslGrade,d.Name)
	if err != nil {
		return err
	}
	return nil
}

