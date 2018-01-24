package strings

import "strings"

// Fetch information from slice of comments (docs).
// Returns appendix of first comment which has tag as prefix.
func FetchMetaInfo(tag string, comments []string) string {
	for _, comment := range comments {
		if len(comment) > len(tag) && strings.HasPrefix(comment, tag) {
			return comment[len(tag)+1:]
		}
	}
	return ""
}