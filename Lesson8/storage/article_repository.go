package storage

import (
	"errors"
	"fmt"
	"go_advanced/Lesson8/internal/app/models"
	"log"
)

type ArticleRepository struct {
	storage *Storage
}

var tableArticle = "articles"

func (ar *ArticleRepository) Create(a *models.Article) (*models.Article, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, author, content) VALUES($1, $2, $3) RETURNING id", tableArticle)
	if err := ar.storage.db.QueryRow(query, a.Title, a.Author, a.Content).Scan(&a.ID); err != nil {

		return nil, err
	}

	return a, nil
}

func (ar *ArticleRepository) DeleteById(id int) (*models.Article, error) {
	article, ok, err := ar.FindArticleById(id)

	if err != nil {

		return nil, err
	}

	if ok {
		query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tableArticle)
		_, err = ar.storage.db.Exec(query, id)

		if err != nil {

			return nil, err
		}
	}

	return article, nil
}

func (ar *ArticleRepository) FindArticleById(id int) (*models.Article, bool, error) {
	var founded bool

	query := fmt.Sprintf("SELECT id, title, author, content FROM %s WHERE id = %d", tableArticle, id)

	rows, err := ar.storage.db.Query(query)
	if err != nil {

		return nil, false, err
	}

	defer rows.Close()

	a := models.Article{}

	for rows.Next() {
		err = rows.Scan(&a.ID, &a.Title, &a.Author, &a.Content)
		if err != nil {

			return nil, false, err
		}

		founded = true
	}

	return &a, founded, nil
}

func (ar *ArticleRepository) UpdateArticle(a *models.Article, id int) (int, bool, error) {
	article, ok, err := ar.FindArticleById(id)
	if err != nil {
		return 0, true, err
	}

	if !ok {
		return 0, false, errors.New("Article not found")
	}

	query := "UPDATE articles SET title = $1, author = $2, content = $3 WHERE id = $4"
	_, err = ar.storage.db.Exec(query, a.Title, a.Author, a.Content, article.ID)
	if err != nil {

		return 0, true, err
	}

	return article.ID, true, nil
}

func (ar *ArticleRepository) SelectAll() ([]*models.Article, error) {
	query := fmt.Sprintf("SELECT* FROM %s", tableArticle)

	rows, err := ar.storage.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	articles := make([]*models.Article, 0, 20)

	for rows.Next() {
		a := models.Article{}

		err = rows.Scan(&a.ID, &a.Title, &a.Author, &a.Content)

		if err != nil {
			log.Println(err)
			continue
		}

		articles = append(articles, &a)
	}

	return articles, nil
}
