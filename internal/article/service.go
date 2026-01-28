package article

func Summarize(url string) (*SummarizeResponse, error) {
	return &SummarizeResponse{
		Title:   "더미 기사",
		Summary: []string{"요약 1", "요약 2", "요약 3"},
	}, nil
}
