package news

import "time"

type (
	News struct {
		ID          int
		Serial      string
		TopicSerial string
		Status      string
		Title       string
		AuthorName  string
		Description string
		CreatedAt   time.Time
		UpdatedAt   time.Time
		DeletedAt   time.Time
	}

	NewsTag struct {
		ID         int
		NewsSerial string
		TagName    string
		CreatedAt  time.Time
		UpdatedAt  time.Time
		DeletedAt  time.Time
	}
)

var (
	NewsTableName    = "news"
	NewsTagTableName = "news_tag"
)
