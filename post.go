package reddit

import (
	"fmt"
	"net/http"
	"strings"
)

type (
	Post struct {
		ApprovedAtUtc              any     `json:"approved_at_utc,omitempty"`
		Subreddit                  string  `json:"subreddit,omitempty"`
		Selftext                   string  `json:"selftext,omitempty"`
		AuthorFullname             string  `json:"author_fullname,omitempty"`
		Saved                      bool    `json:"saved,omitempty"`
		ModReasonTitle             any     `json:"mod_reason_title,omitempty"`
		Gilded                     int     `json:"gilded,omitempty"`
		Clicked                    bool    `json:"clicked,omitempty"`
		IsGallery                  bool    `json:"is_gallery,omitempty"`
		Title                      string  `json:"title,omitempty"`
		LinkFlairRichtext          []any   `json:"link_flair_richtext,omitempty"`
		SubredditNamePrefixed      string  `json:"subreddit_name_prefixed,omitempty"`
		Hidden                     bool    `json:"hidden,omitempty"`
		Pwls                       any     `json:"pwls,omitempty"`
		LinkFlairCSSClass          any     `json:"link_flair_css_class,omitempty"`
		Downs                      int     `json:"downs,omitempty"`
		ThumbnailHeight            int     `json:"thumbnail_height,omitempty"`
		TopAwardedType             any     `json:"top_awarded_type,omitempty"`
		HideScore                  bool    `json:"hide_score,omitempty"`
		Name                       string  `json:"name,omitempty"`
		Quarantine                 bool    `json:"quarantine,omitempty"`
		LinkFlairTextColor         string  `json:"link_flair_text_color,omitempty"`
		UpvoteRatio                float64 `json:"upvote_ratio,omitempty"`
		AuthorFlairBackgroundColor any     `json:"author_flair_background_color,omitempty"`
		SubredditType              string  `json:"subreddit_type,omitempty"`
		Ups                        int     `json:"ups,omitempty"`
		TotalAwardsReceived        int     `json:"total_awards_received,omitempty"`
		MediaEmbed                 struct {
			Content   string `json:"content,omitempty"`
			Width     int    `json:"width,omitempty"`
			Scrolling bool   `json:"scrolling,omitempty"`
			Height    int    `json:"height,omitempty"`
		} `json:"media_embed,omitempty"`
		ThumbnailWidth        int   `json:"thumbnail_width,omitempty"`
		AuthorFlairTemplateID any   `json:"author_flair_template_id,omitempty"`
		IsOriginalContent     bool  `json:"is_original_content,omitempty"`
		UserReports           []any `json:"user_reports,omitempty"`
		SecureMedia           struct {
			RedditVideo struct {
				BitrateKbps       int    `json:"bitrate_kbps"`
				FallbackURL       string `json:"fallback_url"`
				Height            int    `json:"height"`
				Width             int    `json:"width"`
				ScrubberMediaURL  string `json:"scrubber_media_url"`
				DashURL           string `json:"dash_url"`
				Duration          int    `json:"duration"`
				HlsURL            string `json:"hls_url"`
				IsGif             bool   `json:"is_gif"`
				TranscodingStatus string `json:"transcoding_status"`
			} `json:"reddit_video"`
		} `json:"secure_media,omitempty"`
		IsRedditMediaDomain bool `json:"is_reddit_media_domain,omitempty"`
		IsMeta              bool `json:"is_meta,omitempty"`
		Category            any  `json:"category,omitempty"`
		SecureMediaEmbed    struct {
		} `json:"secure_media_embed,omitempty"`
		LinkFlairText       any    `json:"link_flair_text,omitempty"`
		CanModPost          bool   `json:"can_mod_post,omitempty"`
		Score               int    `json:"score,omitempty"`
		ApprovedBy          any    `json:"approved_by,omitempty"`
		IsCreatedFromAdsUI  bool   `json:"is_created_from_ads_ui,omitempty"`
		AuthorPremium       bool   `json:"author_premium,omitempty"`
		Thumbnail           string `json:"thumbnail,omitempty"`
		Edited              any    `json:"edited,omitempty"`
		AuthorFlairCSSClass any    `json:"author_flair_css_class,omitempty"`
		AuthorFlairRichtext []any  `json:"author_flair_richtext,omitempty"`
		Gildings            struct {
		} `json:"gildings,omitempty"`
		PostHint                 string                          `json:"post_hint,omitempty"`
		ContentCategories        any                             `json:"content_categories,omitempty"`
		IsSelf                   bool                            `json:"is_self,omitempty"`
		ModNote                  any                             `json:"mod_note,omitempty"`
		Created                  float64                         `json:"created,omitempty"`
		LinkFlairType            string                          `json:"link_flair_type,omitempty"`
		Wls                      any                             `json:"wls,omitempty"`
		RemovedByCategory        string                          `json:"removed_by_category,omitempty"`
		BannedBy                 any                             `json:"banned_by,omitempty"`
		AuthorFlairType          string                          `json:"author_flair_type,omitempty"`
		Domain                   string                          `json:"domain,omitempty"`
		AllowLiveComments        bool                            `json:"allow_live_comments,omitempty"`
		SelftextHTML             any                             `json:"selftext_html,omitempty"`
		Likes                    any                             `json:"likes,omitempty"`
		SuggestedSort            string                          `json:"suggested_sort,omitempty"`
		BannedAtUtc              any                             `json:"banned_at_utc,omitempty"`
		URLOverriddenByDest      string                          `json:"url_overridden_by_dest,omitempty"`
		ViewCount                any                             `json:"view_count,omitempty"`
		Archived                 bool                            `json:"archived,omitempty"`
		NoFollow                 bool                            `json:"no_follow,omitempty"`
		IsCrosspostable          bool                            `json:"is_crosspostable,omitempty"`
		Pinned                   bool                            `json:"pinned,omitempty"`
		Over18                   bool                            `json:"over_18,omitempty"`
		Preview                  PostPreview                     `json:"preview,omitempty"`
		AllAwardings             []any                           `json:"all_awardings,omitempty"`
		Awarders                 []any                           `json:"awarders,omitempty"`
		MediaOnly                bool                            `json:"media_only,omitempty"`
		CanGild                  bool                            `json:"can_gild,omitempty"`
		Spoiler                  bool                            `json:"spoiler,omitempty"`
		Locked                   bool                            `json:"locked,omitempty"`
		AuthorFlairText          any                             `json:"author_flair_text,omitempty"`
		TreatmentTags            []any                           `json:"treatment_tags,omitempty"`
		Visited                  bool                            `json:"visited,omitempty"`
		RemovedBy                any                             `json:"removed_by,omitempty"`
		NumReports               any                             `json:"num_reports,omitempty"`
		Distinguished            any                             `json:"distinguished,omitempty"`
		SubredditID              string                          `json:"subreddit_id,omitempty"`
		AuthorIsBlocked          bool                            `json:"author_is_blocked,omitempty"`
		ModReasonBy              any                             `json:"mod_reason_by,omitempty"`
		RemovalReason            any                             `json:"removal_reason,omitempty"`
		LinkFlairBackgroundColor string                          `json:"link_flair_background_color,omitempty"`
		ID                       string                          `json:"id,omitempty"`
		IsRobotIndexable         bool                            `json:"is_robot_indexable,omitempty"`
		ReportReasons            any                             `json:"report_reasons,omitempty"`
		Author                   string                          `json:"author,omitempty"`
		DiscussionType           any                             `json:"discussion_type,omitempty"`
		NumComments              int                             `json:"num_comments,omitempty"`
		SendReplies              bool                            `json:"send_replies,omitempty"`
		WhitelistStatus          any                             `json:"whitelist_status,omitempty"`
		ContestMode              bool                            `json:"contest_mode,omitempty"`
		ModReports               []any                           `json:"mod_reports,omitempty"`
		AuthorPatreonFlair       bool                            `json:"author_patreon_flair,omitempty"`
		AuthorFlairTextColor     any                             `json:"author_flair_text_color,omitempty"`
		Permalink                string                          `json:"permalink,omitempty"`
		ParentWhitelistStatus    any                             `json:"parent_whitelist_status,omitempty"`
		Stickied                 bool                            `json:"stickied,omitempty"`
		URL                      string                          `json:"url,omitempty"`
		SubredditSubscribers     int                             `json:"subreddit_subscribers,omitempty"`
		CreatedUtc               float64                         `json:"created_utc,omitempty"`
		NumCrossposts            int                             `json:"num_crossposts,omitempty"`
		Media                    any                             `json:"media,omitempty"`
		MediaMetadata            map[string]PostGalleryMediaData `json:"media_metadata"`
		GalleryData              PostGalleryData                 `json:"gallery_data"`
		IsVideo                  bool                            `json:"is_video,omitempty"`
	}

	PostPreview struct {
		Images             []PostPreviewImage     `json:"images,omitempty"`
		RedditVideoPreview PostPreviewRedditVideo `json:"reddit_video_preview"`
		Enabled            bool                   `json:"enabled,omitempty"`
	}

	PostPreviewImage struct {
		Source      PostPreviewImageSource       `json:"source,omitempty"`
		Resolutions []PostPreviewImageResolution `json:"resolutions,omitempty"`
		Variants    PostPreviewImageVariant      `json:"variants,omitempty"`
		ID          string                       `json:"id,omitempty"`
	}

	PostPreviewImageSource struct {
		URL    string `json:"url,omitempty"`
		Width  int    `json:"width,omitempty"`
		Height int    `json:"height,omitempty"`
	}

	PostPreviewImageResolution struct {
		URL    string `json:"url,omitempty"`
		Width  int    `json:"width,omitempty"`
		Height int    `json:"height,omitempty"`
	}

	PostPreviewImageVariant struct {
		Obfuscated struct {
			Source      PostPreviewImageSource       `json:"source,omitempty"`
			Resolutions []PostPreviewImageResolution `json:"resolutions,omitempty"`
		} `json:"obfuscated,omitempty"`
		Gif struct {
			Source      PostPreviewImageSource       `json:"source,omitempty"`
			Resolutions []PostPreviewImageResolution `json:"resolutions,omitempty"`
		} `json:"gif,omitempty"`
		Mp4 struct {
			Source      PostPreviewImageSource       `json:"source,omitempty"`
			Resolutions []PostPreviewImageResolution `json:"resolutions,omitempty"`
		} `json:"mp4,omitempty"`
		Nsfw struct {
			Source      PostPreviewImageSource       `json:"source,omitempty"`
			Resolutions []PostPreviewImageResolution `json:"resolutions,omitempty"`
		} `json:"nsfw,omitempty"`
	}

	PostGalleryMediaData struct {
		Status string `json:"status"`
		E      string `json:"e"`
		M      string `json:"m"`
		O      []struct {
			Y int    `json:"y"`
			X int    `json:"x"`
			U string `json:"u"`
		} `json:"o"`
		P []struct {
			Y int    `json:"y"`
			X int    `json:"x"`
			U string `json:"u"`
		} `json:"p"`
		S struct {
			Y int    `json:"y"`
			X int    `json:"x"`
			U string `json:"u"`
		} `json:"s"`
		ID string `json:"id"`
	}

	PostGalleryData struct {
		Items []PostGalleryMediaItem `json:"items"`
	}

	PostGalleryMediaItem struct {
		MediaID string `json:"media_id"`
		ID      int    `json:"id"`
	}

	PostPreviewRedditVideo struct {
		BitrateKbps       int    `json:"bitrate_kbps"`
		FallbackURL       string `json:"fallback_url"`
		Height            int    `json:"height"`
		Width             int    `json:"width"`
		ScrubberMediaURL  string `json:"scrubber_media_url"`
		DashURL           string `json:"dash_url"`
		Duration          int    `json:"duration"`
		HlsURL            string `json:"hls_url"`
		IsGif             bool   `json:"is_gif"`
		TranscodingStatus string `json:"transcoding_status"`
	}

	PostHint string
	PostKind string
)

const (
	PostHintImage       PostHint = "image"
	PostHintGif         PostHint = "animatedimage"
	PostHintLink        PostHint = "link"
	PostHintRichVideo   PostHint = "rich:video"
	PostHintRedditVideo PostHint = "hosted:video"
	PostHintGallery     PostHint = ""

	PostKindComment   PostKind = "t1"
	PostKindAccount   PostKind = "t2"
	PostKindLink      PostKind = "t3"
	PostKindMessage   PostKind = "t4"
	PostKindSubreddit PostKind = "t5"
	PostKindAward     PostKind = "t6"
)

var (
	postRedditURL    = "https://www.reddit.com/user/%s/comments/%s.json"
	postSubRedditURL = "https://www.reddit.com/r/%s/comments/%s.json"
)

func (r Reddit) GetPost(subName string, postID string, fromUser bool) ([]RedditJSON, error) {
	url := fmt.Sprintf(postSubRedditURL, subName, postID)
	if fromUser {
		url = fmt.Sprintf(postRedditURL, subName, postID)
	}

	posts := make([]RedditJSON, 0)
	resp, err := r.http.R().SetResult(&posts).Get(url)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode() {
	case http.StatusNotFound:
		return nil, ErrNotFound
	case http.StatusForbidden:
		if strings.Contains(string(resp.Body()), "You've been blocked by network security.") {
			return nil, ErrIPBlocked
		}
		return nil, ErrPrivateSubreddit
	}

	if len(posts) < 1 {
		return nil, ErrNotFound
	}

	if posts[0].Data.Children[0].Data.RemovedByCategory == "deleted" {
		return nil, ErrNotFound
	}
	return posts, nil
}
