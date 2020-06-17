package database

import (
	"database/sql"
	"log"
	"time"
)

var DOMAIN_TABLE string = "domain"

type DomainDBRepository struct {
	Name 				string 		`json:"name,omitempty"`
	SslGrade 			string 		`json:"ssl_grade,omitempty"`
	PreviousSslGrade 	string 		`json:"previous_ssl_grade,omitempty"`
	SearchedAt 			*time.Time 	`json:"-"`
}

type domainRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) domainRepository {
	return domainRepository{db: db, tracer: tracer}
}

func (r domainRepository) CreateDomain(d *DomainDBRepository) error {
	sqlStm := `INSERT INTO `+DOMAIN_TABLE+` (name, ssl_grade, previous_ssl_grade, searched_at) 
	VALUES ($1, $2, $3, NOW())`
	_, err := r.db.Exec(sqlStm, d.Name, d.SslGrade, d.PreviousSslGrade, d.SearchedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r domainRepository) FetchDomains() ([]DomainDBRepository, error) {
	sqlStm := `SELECT * FROM `+DOMAIN_TABLE
	rows, err := r.db.Query(sqlStm)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var domains []DomainDBRepository

	for rows.Next() {
		var d DomainDBRepository
		if err := rows.Scan(&d.Name, &d.SslGrade, &s.PreviousSslGrade, &d.SearchedAt); err != nil {
			log.Println(err)
			continue
		}
		domains = append(domains, d)
	}
	return domains, nil
}

func (r domainRepository) UpdateDomain(d DomainDBRepository) error {
	sqlStm := `UPDATE `+DOMAIN_TABLE+` SET SslGrade = ? , PreviousSslGrade = ? , SearchedAt = NOW() WHERE Name = ?`
	_, err := r.db.Exec(sqlStm,d.SslGrade,d.PreviousSslGrade,d.Name)
	if err != nil {
		return err
	}
	return nil
}

