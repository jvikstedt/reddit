package reddit

import (
	"encoding/json"
	"net/http"
)

func FetchSubReddit(subReddit string) (SubReddit, int) {
	var data SubReddit

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.reddit.com/r/"+subReddit+".json", nil)
	req.Header.Set("User-Agent", "Redmit")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err == nil {
		json.NewDecoder(resp.Body).Decode(&data)
	}
	return data, resp.StatusCode
}

func FetchPost(permalink string) []Post {
	var data []Post

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.reddit.com"+permalink+".json", nil)
	req.Header.Set("User-Agent", "Redmit")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err == nil {
		json.NewDecoder(resp.Body).Decode(&data)
	}

	return data
}
