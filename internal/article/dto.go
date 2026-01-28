package article

type SummarizeRequest struct {
	URL string `json:"url"`
}

type SummarizeResponse struct {
	Title   string   `json:"title"`
	Summary []string `json:"summary"`
}
