package oddsapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	oddsApiHost    = "api.the-odds-api.com"
	sportsEndpoint = "/v4/sports"
)

type Client struct {
	ApiKey            string
	RequestsRemaining int
	RequestsUsed      int
}

func NewClient(apiKey string) (c Client) {
	return Client{
		ApiKey: apiKey,
	}
}

type GetSportsInput struct {
	IncludeInactive bool
}

type GetEventsInput struct {
	Sports     []string
	Regions    []RegionType
	Markets    []MarketType
	OddsFormat OddsFormatType
}

func (c *Client) updateRequestStats(resp *http.Response) {
	h := resp.Header.Get("X-Requests-Remaining")
	r, err := strconv.ParseInt(h, 10, 64)
	if err != nil {
		return
	}
	c.RequestsRemaining = int(r)

	h = resp.Header.Get("X-Requests-Used")
	r, err = strconv.ParseInt(h, 10, 64)
	if err != nil {
		return
	}
	c.RequestsRemaining = int(r)
}

func checkErr(code int) error {
	switch code {
	case http.StatusUnauthorized:
		return ErrUnauthenticated
	case http.StatusUnprocessableEntity:
		return ErrQueryParams
	case http.StatusTooManyRequests:
		return ErrThrottled
	case http.StatusInternalServerError:
		return ErrServer
	default:
		return ErrUnknown
	}
}

// GetSports reaches out to the-odds-api and returns the list of in or out of season sports
func (c *Client) GetSports(input GetSportsInput) (sports []Sport, err error) {

	client := &http.Client{}
	u := &url.URL{
		Host:   oddsApiHost,
		Scheme: "https",
		Path:   fmt.Sprintf("%s/", sportsEndpoint),
	}

	v := url.Values{}
	v.Add("apiKey", c.ApiKey)
	if input.IncludeInactive {
		v.Add("all", "true")
	}
	u.RawQuery = v.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return sports, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return sports, err
	}
	// update request remaining and used
	c.updateRequestStats(resp)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return sports, checkErr(resp.StatusCode)
	}

	// read data
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return sports, err
	}

	err = json.Unmarshal(b, &sports)
	return sports, err
}

// GetEvents reaches out to the-odds-api and returns the latest odds
func (c *Client) GetEvents(input GetEventsInput) (evs []Event, err error) {

	client := &http.Client{}
	v := url.Values{}
	v.Add("apiKey", c.ApiKey)
	for _, r := range input.Regions {
		v.Add("regions", r.String())

	}
	for _, m := range input.Markets {
		v.Add("markets", m.String())

	}

	// odds format
	v.Add("oddsFormat", input.OddsFormat.String())

	for _, s := range input.Sports {
		u := &url.URL{
			Host:   oddsApiHost,
			Scheme: "https",
			Path:   fmt.Sprintf("%s/%s/odds", sportsEndpoint, s),
		}

		u.RawQuery = v.Encode()

		req, err := http.NewRequest("GET", u.String(), nil)
		if err != nil {
			return evs, err
		}

		resp, err := client.Do(req)
		if err != nil {
			return evs, err
		}
		c.updateRequestStats(resp)
		if resp.StatusCode != http.StatusOK {
			return evs, checkErr(resp.StatusCode)
		}

		defer resp.Body.Close()

		// read data
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return evs, err
		}

		var tmpOdds []Event
		err = json.Unmarshal(b, &tmpOdds)
		if err != nil {
			return evs, err
		}

		// convert the raw format to the standard one
		evs = append(evs, tmpOdds...)
	}

	return evs, err
}
