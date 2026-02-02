package news

type NewsItem struct {
	ID       int64  `json:"id"`
	Headline string `json:"headline"`
	Summary  string `json:"summary"`
	Source   string `json:"source"`
	URL      string `json:"url"`
	Image    string `json:"image"`
	Datetime int64  `json:"datetime"`
	Category string `json:"category"`
	Related  string `json:"related,omitempty"`
}

type MarketNewsResponse struct {
	News []NewsItem `json:"news"`
}

type CompanyNewsResponse struct {
	Symbol string     `json:"symbol"`
	News   []NewsItem `json:"news"`
}
