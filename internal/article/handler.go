package article

import (
	"market-brief-be/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

// ArticleHandler HTTP 핸들러
type ArticleHandler struct {
	service ArticleService
}

// NewArticleHandler 생성자
func NewArticleHandler(service ArticleService) *ArticleHandler {
	return &ArticleHandler{service: service}
}

// Register 라우트 등록
func (h *ArticleHandler) Register(r *gin.Engine) {
	r.POST("/articles/summarize", h.Summarize)
}

// Summarize 기사 요약 핸들러
func (h *ArticleHandler) Summarize(c *gin.Context) {
	var req SummarizeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	result, err := h.service.Summarize(req.URL)
	if err != nil {
		response.InternalError(c, "failed to summarize article")
		return
	}

	response.Success(c, result)
}
