package service

import (
	"errors"
	"fmt"
)

type BookmarkService struct{}

var msg string

func (bms *BookmarkService) ViewBookMarks(bm map[string]string) {
	fmt.Println("-----------")
	for key, value := range bm {
		fmt.Printf("Название закладки: %s Содержимое закладки: %s\n", key, value)
	}
	fmt.Println("-----------")
}

func (bms *BookmarkService) AddBookMarks(bm map[string]string, name, text string) (string, error) {

	if _, ok := bm[name]; !ok {
		bm[name] = text
		msg = "Закладка добавлена"

		return msg, nil
	}

	msg = "Закладка с таким названием уже есть"

	return msg, errors.New("there is already a bookmark with this name")
}

func (bms *BookmarkService) DeleteBookMarks(bm map[string]string, name string) (string, error) {
	if _, ok := (bm)[name]; ok {
		delete(bm, name)

		msg = "Закладка удалена"
		return msg, nil
	}

	msg = "Закладки с таким названием нет"

	return msg, errors.New("there is no bookmark with this name")
}
