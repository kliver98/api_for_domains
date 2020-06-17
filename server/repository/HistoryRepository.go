package database

import (
	"database/sql"
	"log"

	database "../database"
)

var HISTORY_TABLE string = "domain"

type IHistoryRepository interface {
	CreateHistory(h *database.HistoryDB) error
	FetchHistories() ([]database.HistoryDB, error)
}

type historyRepository struct {
	db *sql.DB
}

func NewHistoryRepository(db *sql.DB) historyRepository {
	return historyRepository{db: db}
}

func (r historyRepository) CreateHistory(h *database.HistoryDB) error {
	sqlStm := `INSERT INTO `+HISTORY_TABLE+` (domain, searched_at) 
	VALUES ($1, NOW())`
	_, err := r.db.Exec(sqlStm, h.Domain, h.SearchedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r historyRepository) FetchHistories() ([]database.HistoryDB, error) {
	sqlStm := `SELECT * FROM `+HISTORY_TABLE
	rows, err := r.db.Query(sqlStm)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var domains []database.HistoryDB

	for rows.Next() {
		var h database.HistoryDB
		if err := rows.Scan(&h.Domain, &h.SearchedAt); err != nil {
			log.Println(err)
			continue
		}
		domains = append(domains, h)
	}
	return domains, nil
}

