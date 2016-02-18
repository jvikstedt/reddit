package reddit

type SubReddit struct {
	Data struct {
		After    string      `json:"after"`
		Before   interface{} `json:"before"`
		Children []struct {
			Data struct {
				ApprovedBy          interface{}   `json:"approved_by"`
				Archived            bool          `json:"archived"`
				Author              string        `json:"author"`
				AuthorFlairCSSClass string        `json:"author_flair_css_class"`
				AuthorFlairText     string        `json:"author_flair_text"`
				BannedBy            interface{}   `json:"banned_by"`
				Clicked             bool          `json:"clicked"`
				Created             int           `json:"created"`
				CreatedUtc          int           `json:"created_utc"`
				Distinguished       interface{}   `json:"distinguished"`
				Domain              string        `json:"domain"`
				Downs               int           `json:"downs"`
				Edited              int           `json:"edited"`
				From                interface{}   `json:"from"`
				FromID              interface{}   `json:"from_id"`
				FromKind            interface{}   `json:"from_kind"`
				Gilded              int           `json:"gilded"`
				Hidden              bool          `json:"hidden"`
				HideScore           bool          `json:"hide_score"`
				ID                  string        `json:"id"`
				IsSelf              bool          `json:"is_self"`
				Likes               interface{}   `json:"likes"`
				LinkFlairCSSClass   interface{}   `json:"link_flair_css_class"`
				LinkFlairText       interface{}   `json:"link_flair_text"`
				Locked              bool          `json:"locked"`
				Media               interface{}   `json:"media"`
				MediaEmbed          struct{}      `json:"media_embed"`
				ModReports          []interface{} `json:"mod_reports"`
				Name                string        `json:"name"`
				NumComments         int           `json:"num_comments"`
				NumReports          interface{}   `json:"num_reports"`
				Over18              bool          `json:"over_18"`
				Permalink           string        `json:"permalink"`
				Quarantine          bool          `json:"quarantine"`
				RemovalReason       interface{}   `json:"removal_reason"`
				ReportReasons       interface{}   `json:"report_reasons"`
				Saved               bool          `json:"saved"`
				Score               int           `json:"score"`
				SecureMedia         interface{}   `json:"secure_media"`
				SecureMediaEmbed    struct{}      `json:"secure_media_embed"`
				Selftext            string        `json:"selftext"`
				SelftextHTML        string        `json:"selftext_html"`
				Stickied            bool          `json:"stickied"`
				Subreddit           string        `json:"subreddit"`
				SubredditID         string        `json:"subreddit_id"`
				SuggestedSort       interface{}   `json:"suggested_sort"`
				Thumbnail           string        `json:"thumbnail"`
				Title               string        `json:"title"`
				Ups                 int           `json:"ups"`
				URL                 string        `json:"url"`
				UserReports         []interface{} `json:"user_reports"`
				Visited             bool          `json:"visited"`
			} `json:"data"`
			Kind string `json:"kind"`
		} `json:"children"`
		Modhash string `json:"modhash"`
	} `json:"data"`
	Kind string `json:"kind"`
}
