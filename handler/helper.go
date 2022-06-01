package handler

import (
	"fmt"
	"sort"
	"strings"
)

var (
	newsKeyFormat = "%s-%s"
	tagKeysFormat = "%s-%s"
)

func generateNewsCacheKey(status string, topicSerials []string) string {
	sort.Strings(topicSerials)
	return fmt.Sprintf(newsKeyFormat, status, strings.Join(topicSerials, "-"))
}

func generateTagsCacheKey(page, size int) {

}
