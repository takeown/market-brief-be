package router

import (
	"market-brief-be/internal/article"
	"market-brief-be/internal/health"

	"github.com/gin-gonic/gin"
)

func New(articleHandler *article.ArticleHandler) *gin.Engine {
	r := gin.Default()

	health.Register(r)
	articleHandler.Register(r)

	return r
}
