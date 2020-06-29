package database

import (
	"database/sql"

	database "github.com/kliver98/api_for_domains/server/database"
	errors "github.com/kliver98/api_for_domains/server/errors"
)

var HISTORY_TABLE string = "history"

//Methods that must have HistoryRepository
type IHistoryRepository interface {
	CreateHistory(h *database.HistoryDB) error
	FetchHistory() ([]database.HistoryDB, error)
}

type historyRepository struct {
	db *sql.DB
}

func NewHistoryRepository(db *sql.DB) historyRepository {
	return historyRepository{db: db}
}

func (r historyRepository) CreateHistory(newDomain string) error {
	if database.IsDown() {
		return &errors.NoPingError{Message: "Unreachable to database "+database.PING_DATABASE}
	}
	sqlStm := `INSERT INTO `+HISTORY_TABLE+` (domain, searched_at) 
	VALUES ($1, NOW())`
	_, err := r.db.Exec(sqlStm, newDomain)
	if err != nil {
		return &errors.ExecError{Message: "Execution error "+sqlStm+":"+newDomain}
	}
	return nil
}

func (r historyRepository) FetchHistory() ([]database.HistoryDB, error) {
	if database.IsDown() {
		return nil, &errors.NoPingError{Message: "Unreachable to database "+database.PING_DATABASE}
	}
	sqlStm := `SELECT * FROM `+HISTORY_TABLE
	rows, err := r.db.Query(sqlStm)

	if err != nil {
		return nil, &errors.QueryError{Message: err.Error()}
	}

	defer rows.Close()
	var domains []database.HistoryDB
	for rows.Next() {
		var h database.HistoryDB
		if err := rows.Scan(&h.Domain, &h.SearchedAt); err != nil {
			e := &errors.ReadError{}
			h.Domain = e.Error()
			continue
		}
		domains = append(domains, h)
	}
	return domains, nil
}

