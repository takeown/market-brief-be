package article

// ArticleService 인터페이스 정의
type ArticleService interface {
	Summarize(url string) (*SummarizeResponse, error)
}

// articleService 구현체
type articleService struct{}

// NewArticleService 생성자
func NewArticleService() ArticleService {
	return &articleService{}
}

// Summarize 기사 요약
func (s *articleService) Summarize(url string) (*SummarizeResponse, error) {
	// TODO: 실제 스크래핑 및 요약 로직 구현
	return &SummarizeResponse{
		Title:   "더미 기사",
		Summary: []string{"요약 1", "요약 2", "요약 3"},
	}, nil
}
