package service

import (
	"database/sql"

	_ "github.com/lib/pq"

	repository "../repository"
	model "../model"
)

type IHistoryService interface {

	FetchHistory(db *sql.DB) (*model.History, error)
}


func FetchHistories(db *sql.DB) (*model.History,error) { //Here formatted to just return domain names, not times
	historyRepository := repository.NewHistoryRepository(db)
	items, err := historyRepository.FetchHistories()
	if err != nil {
		return nil, err
	}
	var formattedItems []string
	for _,i := range items {
		formattedItems = append(formattedItems, i.Domain)
	}
	return &model.History{Items: formattedItems}, nil
}