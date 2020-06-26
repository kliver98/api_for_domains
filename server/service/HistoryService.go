package service

import (
	"database/sql"
	_ "github.com/lib/pq"
	"reflect"

	repository "../repository"
	errors "../errors"
	model "../model"
)

type IHistoryService interface {

	FetchHistory(db *sql.DB) (*model.History, error)
}


func FetchHistory(db *sql.DB) (*model.History,error) { //Here formatted to just return domain names, not times
	historyRepository := repository.NewHistoryRepository(db)
	items, err := historyRepository.FetchHistory()
	if err != nil {
		e := []interface{}{err, &errors.QueryError{}}
		tmp := reflect.TypeOf(e[0])
		tmp2 := reflect.TypeOf(e[1])
		if tmp==tmp2 {
			return &model.History{Error: model.Error{Type:"Service Unavailable", Message:err.Error(), Code:"503"}}, err
		}
		return &model.History{Error: model.Error{Type:"Internal Server Error", Message:err.Error(), Code:"500"}}, err
	}
	var formattedItems []string
	for _,v := range items {
		formattedItems = append(formattedItems, v.Domain)
	}
	return &model.History{Items: formattedItems}, nil
}