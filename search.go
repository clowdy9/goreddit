package reddit

import (
	"fmt"
	"strings"
)

const (
	searchURL = "https://old.reddit.com/search.json?q=%s&restrict_sr=off&include_over_18=on&sort=relevance&t=all"
)

func (r Reddit) Search(term, afterID string) (*RedditJSON, error) {
	url := fmt.Sprintf(searchURL, strings.ReplaceAll(term, " ", "+"))
	return r.queryJSON(url, afterID)
}
