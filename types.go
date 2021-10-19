package oddsapi

import (
	"time"
)

type Event struct {
	Id           string      `json:"id"`
	SportKey     string      `json:"sport_key"`
	SportTitle   string      `json:"sport_title"`
	CommenceTime time.Time   `json:"commence_time"`
	HomeTeam     string      `json:"home_team"`
	AwayTeam     string      `json:"away_team"`
	Bookmakers   []Bookmaker `json:"bookmakers"`
}

type Bookmaker struct {
	Key        string   `json:"key"`
	Title      string   `json:"title"`
	LastUpdate string   `json:"last_update"`
	Markets    []Market `json:"markets"`
}

type Market struct {
	Key      string    `json:"key"`
	Outcomes []Outcome `json:"outcomes"`
}

type Outcome struct {
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Points float64 `json:"points"`
}

type Sport struct {
	Key          string `json:"key"`
	Active       bool   `json:"active"`
	Group        string `json:"group"`
	Description  string `json:"description"`
	Title        string `json:"title"`
	HasOutrights bool   `json:"has_outrights"`
}
