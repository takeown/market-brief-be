package router

import (
	"market-brief-be/internal/article"
	"market-brief-be/internal/health"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	health.Register(r)
	article.Register(r)

	return r
}
