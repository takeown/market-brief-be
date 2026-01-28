package main

import "github.com/gin-gonic/gin"

type SummarizeRequest struct {
	URL string `json:"url"`
}

type SummarizeResponse struct {
	Title   string   `json:"title"`
	Summary []string `json:"summary"`
}

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.POST("/articles/summarize", func(c *gin.Context) {
		var req SummarizeRequest

		if err := c.BindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, SummarizeResponse{
			Title:   "더미 기사",
			Summary: []string{"요약 1", "요약 2", "요약 3"},
		})
	})

	r.Run(":8080")
}
