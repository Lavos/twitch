package twitch

import (
	"fmt"
	"net/http"
	"net/url"
	"encoding/json"
	"strings"
)

type FollowsResponse struct {
	Follows []Follow `json:"follows"`
}

type Follow struct {
	Channel Channel `json:"channel"`
}

type Channel struct {
	Status string `json:"status"`
	Name string `json:"name"`
	Game string `json:"game"`
}

type StreamsResponse struct {
	Streams []Stream `json:"streams"`
}

type Stream struct {
	Channel Channel `json:"channel"`
}

type TwitchClient struct {
	UserName string
}

func (t *TwitchClient) Follows() ([]string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.twitch.tv/kraken/users/%s/follows/channels", t.UserName), nil)
	req.Header.Add("Accept", "application/vnd.twitchtv.v2+json")

	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%d Status Returned: %s", resp.StatusCode, resp.Status)
	}

	dec := json.NewDecoder(resp.Body)
	var r FollowsResponse

	json_err := dec.Decode(&r)

	if json_err != nil {
		return nil, json_err
	}

	names := make([]string, len(r.Follows))
	for i, f := range r.Follows {
		names[i] = f.Channel.Name
	}

	return names, nil
}

func (t *TwitchClient) Online() ([]string, error) {
	names, err := t.Follows()

	if err != nil {
		return nil, err
	}

	online_url, _ := url.Parse("https://api.twitch.tv/kraken/streams")
	q := online_url.Query()
	q.Add("channel", strings.Join(names, ","))
	online_url.RawQuery = q.Encode()

	req, _ := http.NewRequest("GET", online_url.String(), nil)
	req.Header.Add("Accept", "application/vnd.twitchtv.v2+json")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%d Status Returned: %s", resp.StatusCode, resp.Status)
	}

	dec := json.NewDecoder(resp.Body)
	var r StreamsResponse

	json_err := dec.Decode(&r)

	if json_err != nil {
		return nil, json_err
	}

	online_names := make([]string, len(r.Streams))
	for i, o := range r.Streams {
		online_names[i] = o.Channel.Name
	}

	return online_names, nil
}
