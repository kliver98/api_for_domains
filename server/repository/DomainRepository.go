package database

import (
	"database/sql"
	"log"
	database "../database"
)

var DOMAIN_TABLE string = "domain"

type IDomainRepository interface {
	CreateDomain(domain string, sslGrade string, previouSsslGrade string) error
	FetchDomain(domain string) (database.DomainDB, error)
	UpdateDomain(domain string, sslGrade string, previouSsslGrade string) error
}

type domainRepository struct {
	db *sql.DB
}

func NewDomainRepository(db *sql.DB) domainRepository {
	return domainRepository{db: db}
}

func (r domainRepository) CreateDomain(domain string, sslGrade string, previousSslGrade string) error {
	var d database.DomainDB
	d.Name = domain
	d.SslGrade = sslGrade
	d.PreviousSslGrade = previousSslGrade
	sqlStm := `INSERT INTO `+DOMAIN_TABLE+` (name, ssl_grade , previous_ssl_grade , searched_at ) 
	VALUES ($1, $2, $3, NOW())`
	_, err := r.db.Exec(sqlStm, d.Name, d.SslGrade, d.PreviousSslGrade)
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

func (r domainRepository) UpdateDomain(domain string, sslGrade string, previouSsslGrade string) error {
	var d database.DomainDB
	d.Name = domain
	d.SslGrade = sslGrade
	d.PreviousSslGrade = previouSsslGrade
	sqlStm := `UPDATE `+DOMAIN_TABLE+` SET ssl_grade = $1 , previous_ssl_grade = $2 , searched_at = NOW() WHERE Name = $3`
	_, err := r.db.Exec(sqlStm,d.SslGrade,d.PreviousSslGrade,d.Name)
	if err != nil {
		return err
	}
	return nil
}

