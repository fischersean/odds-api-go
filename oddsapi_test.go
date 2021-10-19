package oddsapi

import (
	"errors"
	"github.com/jarcoal/httpmock"
	"os"
	"testing"
)

func TestNewClient(t *testing.T) {
	key := "this-is-the-key"
	client := NewClient(key)
	if client.ApiKey != key {
		t.Error("Client API Key not set")
	}
}

func TestGetSportsMock(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.the-odds-api.com/v4/sports/",
		httpmock.NewStringResponder(200, `[
  {
    "key": "americanfootball_nfl",
    "active": true,
    "group": "American Football",
    "description": "US Football",
    "title": "NFL",
    "has_outrights": false
  }
]`))

	key := os.Getenv("ODDS_API_KEY")
	c := NewClient(key)

	s, err := c.GetSports(GetSportsInput{
		IncludeInactive: false,
	})
	if err != nil {
		t.Error(err.Error())
	}

	if s[0].Key != "americanfootball_nfl" {
		t.Error("Could not decode response")
	}
}

func TestGetEventsMock(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.the-odds-api.com/v4/sports/americanfootball_nfl/odds",
		httpmock.NewStringResponder(200, `[
  {
    "id": "e912304de2b2ce35b473ce2ecd3d1502",
    "sport_key": "americanfootball_nfl",
    "sport_title": "NFL",
    "commence_time": "2020-01-02T23:10:00Z",
    "home_team": "Houston Texans",
    "away_team": "Kansas City Chiefs",
    "bookmakers": [
      {
        "key": "draftkings",
        "title": "DraftKings",
        "last_update": "2020-01-01T14:45:13Z",
        "markets": [
          {
            "key": "h2h",
            "outcomes": [
              {
                "name": "Houston Texans",
                "price": 2.23
              },
              {
                "name": "Kansas City Chiefs",
                "price": 1.45
              }
            ]
          }
        ]
      }
    ]
  }
]`))

	key := os.Getenv("ODDS_API_KEY")
	c := NewClient(key)

	s, err := c.GetEvents(GetEventsInput{
		Sports:  []string{"americanfootball_nfl"},
		Regions: []RegionType{RegionUS},
		Markets: []MarketType{MarketMoneyline},
	})
	if err != nil {
		t.Error(err.Error())
	}

	if s[0].Id != "e912304de2b2ce35b473ce2ecd3d1502" {
		t.Error("Could not decode response")
	}

}

func TestErrors(t *testing.T) {
	if !errors.Is(checkErr(401), ErrUnauthenticated) {
		t.Error("Status code 401 not handled")
	}
	if !errors.Is(checkErr(422), ErrQueryParams) {
		t.Error("Status code 422 not handled")
	}
	if !errors.Is(checkErr(429), ErrThrottled) {
		t.Error("Status code 429 not handled")
	}
	if !errors.Is(checkErr(500), ErrServer) {
		t.Error("Status code 500 not handled")
	}
	if !errors.Is(checkErr(999), ErrUnknown) {
		t.Error("Unknown status code not handled")
	}
}

// TestGetEvents actually makes a call to the API.
// Normally this is skipped during testing but should occasionally be run to verify library function
func TestGetEvents(t *testing.T) {

	// SKIP
	t.SkipNow()

	key := os.Getenv("ODDS_API_KEY")
	c := NewClient(key)

	s, err := c.GetEvents(GetEventsInput{
		Sports:     []string{"americanfootball_nfl"},
		Regions:    []RegionType{RegionUS},
		Markets:    []MarketType{MarketMoneyline},
		OddsFormat: OddsFormatDecimal,
	})
	if err != nil {
		t.Error(err.Error())
	}

	if s[0].SportKey != "americanfootball_nfl" {
		t.Errorf("Sport key %s != americanfootball_nfl", s[0].SportKey)
	}

}
