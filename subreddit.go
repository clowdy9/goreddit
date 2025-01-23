package reddit

import (
	"fmt"
)

const (
	subredditURL = "https://www.reddit.com/r/%s/.json"
)

func (r Reddit) Subreddit(name, afterID string) (*RedditJSON, error) {
	url := fmt.Sprintf(subredditURL, name)
	return r.queryJSON(url, afterID)
}
