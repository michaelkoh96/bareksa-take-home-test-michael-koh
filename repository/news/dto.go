package news

type (
	News struct {
		ID          int
		Serial      string
		TopicSerial string
		Status      string
		Title       string
		AuthorName  string
		Description string
	}

	NewsTag struct {
		ID         int
		NewsSerial string
		TagName    string
	}
)

var (
	NewsTableName    = "news"
	NewsTagTableName = "news_tag"
)
