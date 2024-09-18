package handlers

import (
	"github.com/gin-gonic/gin"
	"go_advanced/Lesson10/helpers"
	"go_advanced/Lesson10/models"
)

func GetAllArticles(c *gin.Context) {
	var articles []models.Article
	err := models.GetAllArticles(&articles)
	if err != nil {
		helpers.RespondJSON(c, 404, articles)

		return
	}

	helpers.RespondJSON(c, 200, articles)
}

func PostNewArticle(c *gin.Context) {
	var article models.Article
	c.BindJSON(&article) // Decoding
	err := models.AddNewArticle(&article)
	if err != nil {
		helpers.RespondJSON(c, 404, article)

		return
	}

	helpers.RespondJSON(c, 201, article)
}

func GetArticleById(c *gin.Context) {
	id := c.Param("id")
	var article models.Article
	err := models.GetArticleById(&article, id)

	if err != nil {
		helpers.RespondJSON(c, 404, article)

		return
	}

	helpers.RespondJSON(c, 200, article)
}

func UpdateArticleById(c *gin.Context) {
	id := c.Param("id")
	var article models.Article
	err := models.GetArticleById(&article, id)

	if err != nil {
		helpers.RespondJSON(c, 404, article)

		return
	}

	c.BindJSON(&article)
	err = models.UpdateArticleById(&article, id)
	if err != nil {
		helpers.RespondJSON(c, 404, article)

		return
	}

	helpers.RespondJSON(c, 202, article)
}

func DeleteArticleById(c *gin.Context) {
	id := c.Param("id")
	var article models.Article
	err := models.DeleteArticleById(&article, id)

	if err != nil {
		helpers.RespondJSON(c, 404, article)

		return
	}

	helpers.RespondJSON(c, 202, article)
}
