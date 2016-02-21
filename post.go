package reddit

type Post struct {
	Data struct {
		Children []struct {
			Data struct {
				Replies Post   `json:"replies"`
				Body    string `json:"body"`
				Author  string `json:"author"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}
