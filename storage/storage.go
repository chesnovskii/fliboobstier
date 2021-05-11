package storage

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/chesnovsky/fliboobstier/bot_helpers"
	"github.com/chesnovsky/fliboobstier/logger"
	_ "github.com/mattn/go-sqlite3"
)

type RegexActionElements struct {
	Images    []string
	Stickers  []string
	Gifs      []string
	Documents []string
}

type StorageInstance struct {
	url string
	db  *sql.DB
}

func (storage StorageInstance) GetAdminList() ([]string, error) {
	return storage.GetAdminListSqlite()
}

func (storage StorageInstance) RemoveRegexActionElement(action_id string, element_type string, element_id string) error {
	return storage.RemoveRegexActionElementSqlite(action_id, element_type, element_id)
}

func (storage StorageInstance) AddRegexActionElement(action_id string, element_type string, element_id string) error {
	return storage.AddRegexActionElementSqlite(action_id, element_type, element_id)
}

func (storage StorageInstance) GetRegexActionElements(action_id string) (RegexActionElements, error) {
	err_list := []string{}
	Images, err := storage.GetRegexActionElementsSqLite(action_id, "image")
	if err != nil && err != errors.New("") {
		err_msg := fmt.Sprintf("Error getting regex_action images: %+v", err)
		logger.Logger.Error(err_msg)
		err_list = append(err_list, err_msg)
	}

	Stickers, err := storage.GetRegexActionElementsSqLite(action_id, "sticker")
	if err != nil && err != errors.New("") {
		err_msg := fmt.Sprintf("Error getting regex_action stickers: %+v", err)
		logger.Logger.Error(err_msg)
		err_list = append(err_list, err_msg)
	}

	Documents, err := storage.GetRegexActionElementsSqLite(action_id, "document")
	if err != nil && err != errors.New("") {
		err_msg := fmt.Sprintf("Error getting regex_action documents: %+v", err)
		logger.Logger.Error(err_msg)
		err_list = append(err_list, err_msg)
	}

	Gifs, err := storage.GetRegexActionElementsSqLite(action_id, "gif")
	if err != nil && err != errors.New("") {
		err_msg := fmt.Sprintf("Error getting regex_action gifs: %+v", err)
		logger.Logger.Error(err_msg)
		err_list = append(err_list, err_msg)
	}

	action_elements := RegexActionElements{Images, Stickers, Gifs, Documents}

	return action_elements, bot_helpers.ErrListToError(err_list)
}
