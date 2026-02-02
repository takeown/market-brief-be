package news

import (
	"context"
	"time"

	"market-brief-be/internal/finnhub"
)

type Service struct {
	finnhubClient *finnhub.Client
}

func NewService(finnhubClient *finnhub.Client) *Service {
	return &Service{
		finnhubClient: finnhubClient,
	}
}

func (s *Service) GetMarketNews(ctx context.Context, category string) (*MarketNewsResponse, error) {
	if category == "" {
		category = "general"
	}

	marketNews, err := s.finnhubClient.GetMarketNews(ctx, category)
	if err != nil {
		return nil, err
	}

	items := make([]NewsItem, len(marketNews))
	for i, n := range marketNews {
		items[i] = NewsItem{
			ID:       n.ID,
			Headline: n.Headline,
			Summary:  n.Summary,
			Source:   n.Source,
			URL:      n.URL,
			Image:    n.Image,
			Datetime: n.Datetime,
			Category: n.Category,
			Related:  n.Related,
		}
	}

	return &MarketNewsResponse{News: items}, nil
}

func (s *Service) GetCompanyNews(ctx context.Context, symbol string) (*CompanyNewsResponse, error) {
	now := time.Now()
	to := now.Format("2006-01-02")
	from := now.AddDate(0, 0, -7).Format("2006-01-02")

	companyNews, err := s.finnhubClient.GetCompanyNews(ctx, symbol, from, to)
	if err != nil {
		return nil, err
	}

	items := make([]NewsItem, len(companyNews))
	for i, n := range companyNews {
		items[i] = NewsItem{
			ID:       n.ID,
			Headline: n.Headline,
			Summary:  n.Summary,
			Source:   n.Source,
			URL:      n.URL,
			Image:    n.Image,
			Datetime: n.Datetime,
			Category: n.Category,
			Related:  n.Related,
		}
	}

	return &CompanyNewsResponse{
		Symbol: symbol,
		News:   items,
	}, nil
}
