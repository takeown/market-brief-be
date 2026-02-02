package news

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(r *gin.Engine) {
	group := r.Group("/news")
	{
		group.GET("/market", h.GetMarketNews)
		group.GET("/company/:symbol", h.GetCompanyNews)
	}
}

func (h *Handler) GetMarketNews(c *gin.Context) {
	category := c.DefaultQuery("category", "general")

	resp, err := h.service.GetMarketNews(c.Request.Context(), category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetCompanyNews(c *gin.Context) {
	symbol := c.Param("symbol")
	if symbol == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "symbol is required"})
		return
	}

	resp, err := h.service.GetCompanyNews(c.Request.Context(), symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
