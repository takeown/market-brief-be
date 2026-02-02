package main

import (
	"market-brief-be/internal/article"
	"market-brief-be/internal/config"
	"market-brief-be/internal/finnhub"
	"market-brief-be/internal/news"
	"market-brief-be/internal/router"
)

func main() {
	// 설정 로드
	cfg := config.Load()

	// Finnhub 클라이언트 생성
	finnhubClient := finnhub.NewClient(cfg.FinnhubBaseURL, cfg.FinnhubAPIKey)

	// 서비스 생성
	articleService := article.NewArticleService()
	newsService := news.NewService(finnhubClient)

	// 핸들러 생성
	articleHandler := article.NewArticleHandler(articleService)
	newsHandler := news.NewHandler(newsService)

	// 라우터 구성 및 실행
	r := router.New(articleHandler, newsHandler)
	r.Run(":" + cfg.Port)
}
