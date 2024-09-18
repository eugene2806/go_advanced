package models

import (
	_ "github.com/lib/pq"
	"go_advanced/Lesson10/storage"
)

func GetAllArticles(a *[]Article) error {
	if err := storage.DB.Find(a).Error; err != nil {

		return err
	}

	return nil
}

func AddNewArticle(a *Article) error {

	return storage.DB.Create(a).Error
}

func GetArticleById(a *Article, id string) error {

	return storage.DB.Where("id = ?", id).First(a).Error
}

func DeleteArticleById(a *Article, id string) error {

	return storage.DB.Where("id = ?", id).Delete(a).Error
}

func UpdateArticleById(a *Article, id string) error {

	return storage.DB.Where("id = ?", id).Save(a).Error
}
