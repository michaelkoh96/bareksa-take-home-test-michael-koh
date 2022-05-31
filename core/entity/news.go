package entity

type (
	News struct {
		Serial      string
		TopicSerial string
		Status      string
		Title       string
		AuthorName  string
		Description string
		Tags        []string
	}

	GetNewsQuery struct {
		Status       string
		TopicSerials []string
	}
)
