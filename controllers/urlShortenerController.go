package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thgmagno/api-go/requests"
	"github.com/thgmagno/api-go/services"
)

func ShortenUrl(c *gin.Context) {
	var req requests.ShortenUrlRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "URL inválida"})
		return
	}

	urlShortener, err := services.ShortenUrl(req.URL)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "erro",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "sucesso",
		"url":    urlShortener,
	})
}

func RedirectToOriginalUrl(c *gin.Context) {
	url := c.Query("url")

	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "erro",
			"message": "Parâmetro 'url' não informado.",
			"details": "O parâmetro 'url' é obrigatório.",
		})
		return
	}

	originalUrl, err := services.Redis.Get(services.Ctx, "short:"+url).Result()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  err.Error(),
			"message": "URL não encontrada.",
			"details": "Este projeto é parte de um portfólio pessoal e não possui fins comerciais. As URLs expiram em 30 dias. Caso a URL tenha expirado, gere uma nova.",
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, originalUrl)
}
