package response

type (
	GeneralResponse struct {
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}

	Topic struct {
		Serial string `json:"serial"`
		Title  string `json:"title"`
	}

	News struct {
		Serial      string   `json:"serial"`
		Topic       Topic    `json:"topic"`
		Status      string   `json:"status"`
		Title       string   `json:"title"`
		AuthorName  string   `json:"authorName"`
		Description string   `json:"description"`
		Tags        []string `json:"tag"`
	}

	Tag struct {
		Name string `json:"name"`
	}
)
