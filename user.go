package reddit

import (
	"fmt"
)

const (
	userRedditURL = "https://www.reddit.com/user/%s/.json"
)

func (r Reddit) User(name, afterID string) (*RedditJSON, error) {
	url := fmt.Sprintf(userRedditURL, name)
	return r.queryJSON(url, afterID)
}
