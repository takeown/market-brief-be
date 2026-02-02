package router

import (
	"market-brief-be/internal/article"
	"market-brief-be/internal/health"
	"market-brief-be/internal/news"

	"github.com/gin-gonic/gin"
)

func New(articleHandler *article.ArticleHandler, newsHandler *news.Handler) *gin.Engine {
	r := gin.Default()

	health.Register(r)
	articleHandler.Register(r)
	newsHandler.Register(r)

	return r
}
