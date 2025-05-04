package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

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

	shortUrl, err := services.ShortenUrl(req.URL)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "erro",
			"message": err.Error(),
		})
		return
	}

	urlData := map[string]string{
		"short":    shortUrl,
		"original": req.URL,
	}

	dataJson, err := json.Marshal(urlData)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "erro",
			"message": "Erro ao processar dados",
		})
	}

	services.Redis.LPush(services.Ctx, "urls_shortened", dataJson)

	c.JSON(200, gin.H{
		"status": "sucesso",
		"url":    shortUrl,
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

func RecentlyShortenedUrls(c *gin.Context) {
	takeParam := c.DefaultQuery("take", "10")
	take, err := strconv.ParseInt(takeParam, 10, 64)
	if err != nil || take <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "erro",
			"message": "Parâmetro 'take' inválido",
		})
		return
	}

	urls, err := services.Redis.LRange(services.Ctx, "urls_shortened", 0, take-1).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "erro",
			"message": "Erro ao buscar URLs recentes",
		})
		return
	}

	total, err := services.Redis.LLen(services.Ctx, "urls_shortened").Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "erro",
			"message": "Erro ao verificar total de URLs",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "sucesso",
		"message": "",
		"urls":    urls,
		"total":   total,
	})
}

func FlushAll(c *gin.Context) {
	err := services.Redis.FlushDB(services.Ctx).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "erro",
			"message": "Erro ao limpar o banco Redis.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "sucesso",
		"message": "Todos os dados foram removidos com sucesso.",
	})
}
