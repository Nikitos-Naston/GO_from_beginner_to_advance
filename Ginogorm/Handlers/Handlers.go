package handlers

import (
	"Ginogorm/helpers"
	"Ginogorm/models"

	"github.com/gin-gonic/gin"
)

func GetallArticles(c *gin.Context) {
	var arcticles []models.Arcticle
	err := models.GetallArticles(&arcticles)
	if err != nil {
		helpers.RespondJSON(c, 404, arcticles)
		return
	}
	helpers.RespondJSON(c, 200, arcticles)
}

func PostNewArticl(c *gin.Context) {
	var article models.Arcticle
	c.BindJSON(&article)
	err := models.AddNewArticle(&article)

	if err != nil {
		helpers.RespondJSON(c, 404, article)
		return
	}
	helpers.RespondJSON(c, 201, article)
}

func GetArticlebyID(c *gin.Context) {
	id := c.Params.ByName("id")
	var arcticle models.Arcticle
	err := models.GetArticlebyID(&arcticle, id)
	if err != nil {
		helpers.RespondJSON(c, 404, arcticle)
		return
	}
	helpers.RespondJSON(c, 201, arcticle)

}

func UpdateArcticleById(c *gin.Context) {
	id := c.Params.ByName("id")
	var arcticle models.Arcticle
	err := models.UpdateArticlebyID(&arcticle, id)
	if err != nil {
		helpers.RespondJSON(c, 404, arcticle)
		return
	}
	helpers.RespondJSON(c, 201, arcticle)
}

func DeleteArctilebyId(c *gin.Context) {
	id := c.Params.ByName("id")
	var arcticle models.Arcticle
	err := models.DeleteArticlebyID(&arcticle, id)
	if err != nil {
		helpers.RespondJSON(c, 404, arcticle)
		return
	}
	helpers.RespondJSON(c, 202, arcticle)
}
