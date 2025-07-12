package lastfm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	apiURL = "https://ws.audioscrobbler.com/2.0/"
)

type Client struct {
	apiKey   string
	username string
	client   *http.Client
}

type NowPlaying struct {
	Track      string
	Artist     string
	Album      string
	IsPlaying  bool
	LastPlayed time.Time
}

type RecentTracksResponse struct {
	RecentTracks struct {
		Track []struct {
			Name   string `json:"name"`
			Artist struct {
				Text string `json:"#text"`
			} `json:"artist"`
			Album struct {
				Text string `json:"#text"`
			} `json:"album"`
			Attr *struct {
				NowPlaying string `json:"nowplaying"`
			} `json:"@attr,omitempty"`
			Date *struct {
				UTS string `json:"uts"`
			} `json:"date,omitempty"`
		} `json:"track"`
	} `json:"recenttracks"`
}

func NewClient(apiKey, username string) *Client {
	return &Client{
		apiKey:   apiKey,
		username: username,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *Client) GetNowPlaying() (*NowPlaying, error) {
	url := fmt.Sprintf("%s?method=user.getrecenttracks&user=%s&api_key=%s&format=json&limit=1",
		apiURL, c.username, c.apiKey)

	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("lastfm api error: %d", resp.StatusCode)
	}

	var data RecentTracksResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if len(data.RecentTracks.Track) == 0 {
		return nil, nil
	}

	track := data.RecentTracks.Track[0]
	np := &NowPlaying{
		Track:  track.Name,
		Artist: track.Artist.Text,
		Album:  track.Album.Text,
	}

	if track.Attr != nil && track.Attr.NowPlaying == "true" {
		np.IsPlaying = true
		np.LastPlayed = time.Now()
	} else if track.Date != nil {
		np.IsPlaying = false
	}

	return np, nil
}