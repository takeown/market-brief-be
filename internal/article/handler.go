package article

import "github.com/gin-gonic/gin"

func Register(r *gin.Engine) {
	r.POST("/articles/summarize", func(c *gin.Context) {
		var req SummarizeRequest

		if err := c.BindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		result, err := Summarize(req.URL)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to summarize"})
			return
		}

		c.JSON(200, result)
	})
}
