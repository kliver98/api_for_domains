package database

import (
	"database/sql"

	database "github.com/kliver98/api_for_domains/server/database"
	errors "github.com/kliver98/api_for_domains/server/errors"
)

var DOMAIN_TABLE string = "domain"

//Methods that must have DomainRepository
type IDomainRepository interface {
	CreateDomain(domain string, sslGrade string, previouSsslGrade string) error
	FetchDomain(domain string) (database.DomainDB, error)
	UpdateDomain(domain string, sslGrade string, previouSsslGrade string) error
}

type DomainRepository struct {
	db *sql.DB
}

func NewDomainRepository(db *sql.DB) DomainRepository {
	return DomainRepository{db: db}
}

func (r DomainRepository) CreateDomain(domain string, sslGrade string, previousSslGrade string) error {
	if database.IsDown() {
		return &errors.NoPingError{Message: "Unreachable to database "+database.PING_DATABASE}
	}
	var d database.DomainDB
	d.Name = domain
	d.SslGrade = sslGrade
	d.PreviousSslGrade = previousSslGrade
	sqlStm := `INSERT INTO `+DOMAIN_TABLE+` (name, ssl_grade , previous_ssl_grade , searched_at ) 
	VALUES ($1, $2, $3, NOW())`
	_, err := r.db.Exec(sqlStm, d.Name, d.SslGrade, d.PreviousSslGrade)
	if err != nil {
		return &errors.ExecError{Message: err.Error()+" \n Execution error "+sqlStm+":"+domain+","+sslGrade+","+previousSslGrade}
	}
	return nil
}

func (r DomainRepository) FetchDomain(domain string) ([]database.DomainDB, error) {
	if database.IsDown() {
		return *&[]database.DomainDB{}, &errors.NoPingError{Message: "Unreachable to database "+database.PING_DATABASE}
	}
	sqlStm := `SELECT * FROM `+DOMAIN_TABLE+` WHERE name='`+domain+`'`
	rows, err := r.db.Query(sqlStm)
	if err != nil {
		return *&[]database.DomainDB{}, &errors.QueryError{Message: "A Query error has ocurred "+err.Error()}
	}
	defer rows.Close()
	var domains []database.DomainDB

	for rows.Next() {
		var d database.DomainDB
		if err := rows.Scan(&d.Name, &d.SslGrade, &d.PreviousSslGrade, &d.SearchedAt); err != nil {
			e := &errors.ReadError{"Reading error, possible loss of data"}
			d.Name = e.Error()
			continue
		}
		domains = append(domains, d)
	}
	return domains, nil
}

func (r DomainRepository) UpdateDomain(domain string, sslGrade string, previousSslGrade string) error {
	if database.IsDown() {
		return &errors.NoPingError{Message: "Unreachable to database "+database.PING_DATABASE}
	}
	var d database.DomainDB
	d.Name = domain
	d.SslGrade = sslGrade
	d.PreviousSslGrade = previousSslGrade
	sqlStm := `UPDATE `+DOMAIN_TABLE+` SET ssl_grade = $1 , previous_ssl_grade = $2 , searched_at = NOW() WHERE Name = $3`
	_, err := r.db.Exec(sqlStm,d.SslGrade,d.PreviousSslGrade,d.Name)
	if err != nil {
		return &errors.ExecError{Message: err.Error()+" \n Execution error "+sqlStm+":"+domain+","+sslGrade+","+previousSslGrade}
	}
	return nil
}

