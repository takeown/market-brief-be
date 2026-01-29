package main

import (
	"market-brief-be/internal/article"
	"market-brief-be/internal/config"
	"market-brief-be/internal/router"
)

func main() {
	// 설정 로드
	cfg := config.Load()

	// 서비스 생성
	articleService := article.NewArticleService()

	// 핸들러 생성
	articleHandler := article.NewArticleHandler(articleService)

	// 라우터 구성 및 실행
	r := router.New(articleHandler)
	r.Run(":" + cfg.Port)
}
