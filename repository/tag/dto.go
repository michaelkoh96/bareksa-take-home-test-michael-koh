package tag

import "time"

type (
	Tag struct {
		ID        int
		Name      string
		CreatedAt *time.Time
		UpdatedAt *time.Time
		DeletedAt *time.Time
	}
)

var (
	TagTableName = "tag"
)
