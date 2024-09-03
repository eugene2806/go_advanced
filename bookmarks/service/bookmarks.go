package service

import (
	"bookmarks/model"
	"errors"
)

type BookmarkService struct {
}

func (bms *BookmarkService) ViewBookMarks() (map[string]string, error) {
	var db interface{}
	db = model.DB
	if _, ok := db.(map[string]string); !ok {

		return nil, errors.New("error opening the database")
	}

	return model.DB, nil
}

func (bms *BookmarkService) AddBookMarks(name, text string) error {

	if _, ok := model.DB[name]; !ok {
		model.DB[name] = text

		return nil
	}

	return errors.New("there is already a bookmark with this name")
}

func (bms *BookmarkService) DeleteBookMarks(name string) error {
	if _, ok := (model.DB)[name]; ok {
		delete(model.DB, name)

		return nil
	}

	return errors.New("there is no bookmark with this name")
}
