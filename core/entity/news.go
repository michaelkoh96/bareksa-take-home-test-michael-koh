package entity

type (
	News struct {
		Serial      string   `json:"serial"`
		TopicSerial string   `json:"topicSerial"`
		Status      string   `json:"status"`
		Title       string   `json:"title"`
		AuthorName  string   `json:"authorName"`
		Description string   `json:"description"`
		Tags        []string `json:"tags"`
	}

	GetNewsQuery struct {
		Status       string
		TopicSerials []string
	}
)
