package reddit

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"math/rand"

	"github.com/go-resty/resty/v2"
)

type (
	Reddit struct {
		http             *resty.Client
		Proxies          []string
		UserAgents       []string
		RandomHTTPHeader bool
	}

	RedditOpts struct {
		Proxy            string
		Proxies          []string
		UserAgent        string
		UserAgents       []string
		RandomHTTPHeader bool
	}

	RedditJSON struct {
		Kind string         `json:"kind,omitempty"`
		Data RedditJSONData `json:"data,omitempty"`
	}

	RedditJSONData struct {
		After     string                   `json:"after,omitempty"`
		Dist      int                      `json:"dist,omitempty"`
		Modhash   string                   `json:"modhash,omitempty"`
		GeoFilter any                      `json:"geo_filter,omitempty"`
		Children  []RedditJSONDataChildren `json:"children,omitempty"`
		Before    any                      `json:"before,omitempty"`
	}

	RedditJSONDataChildren struct {
		Kind string `json:"kind,omitempty"`
		Data Post   `json:"data,omitempty"`
	}
)

var (
	userAgent           = "Mozilla/5.0 (Macintosh; Intel Mac OS X 14.4; rv:124.0) Gecko/20100101 Firefox/124.0"
	ErrNotFound         = errors.New("[404] subreddit not found")
	ErrPrivateSubreddit = errors.New("[403] subreddit is private")
	ErrIPBlocked        = errors.New("[403] IP blocked")
)

const (
	paginationLimit = "25"
	paginationCount = "25"
)

func New(opts RedditOpts) *Reddit {
	client := resty.New()
	client.SetRetryCount(1)
	client.SetRetryWaitTime(time.Second * 2)
	if opts.Proxy != "" {
		client.SetProxy(opts.Proxy)
	}

	if opts.UserAgent != "" {
		userAgent = opts.UserAgent
	}

	client.SetHeaders(map[string]string{
		"User-Agent":      userAgent,
		"Accept":          "application/json",
		"Accept-Encoding": "identity",
	})

	return &Reddit{
		http:             client,
		RandomHTTPHeader: opts.RandomHTTPHeader,
		UserAgents:       opts.UserAgents,
		Proxies:          opts.Proxies,
	}
}

func (r *Reddit) WithHTTPHeaders(h map[string]string) {
	r.http.SetHeaders(h)
}

func (r *Reddit) WithHTTPUserAgent(u string) {
	r.http.SetHeader("User-Agent", u)
}

func (r *Reddit) WithHTTPProxy(p string) {
	r.http.SetProxy(p)
}

func (r *Reddit) queryJSON(uri, afterID string) (*RedditJSON, error) {
	if afterID != "" {
		u, err := url.Parse(uri)
		if err != nil {
			return nil, err
		}
		query := u.Query()
		query.Set("after", afterID)
		query.Set("limit", paginationLimit)
		query.Set("count", paginationCount)
		u.RawQuery = query.Encode()
		uri = u.String()

	}

	if r.RandomHTTPHeader {
		s := rand.NewSource(time.Now().Unix())
		ran := rand.New(s)
		if len(r.UserAgents) > 0 {
			r.WithHTTPUserAgent(r.UserAgents[ran.Intn(len(r.UserAgents))])
		}

		if len(r.Proxies) > 0 {
			r.WithHTTPProxy(r.Proxies[ran.Intn(len(r.Proxies))])
		}

	}

	resp, err := r.http.R().SetResult(&RedditJSON{}).Get(uri)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode() {
	case http.StatusOK:
		break
	case http.StatusNotFound:
		return nil, ErrNotFound
	case http.StatusForbidden:
		if strings.Contains(string(resp.Body()), "You've been blocked by network security.") {
			return nil, ErrIPBlocked
		}
		return nil, ErrPrivateSubreddit
	default:
		err = fmt.Errorf("unkown error: status: %d, %s", resp.StatusCode(), string(resp.Body()))
	}

	sub := resp.Result().(*RedditJSON)
	return sub, err
}
