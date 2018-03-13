package main

type Player struct {
	AccountID       int         `json:"account_id"`
	Steamid         interface{} `json:"steamid"`
	Avatar          interface{} `json:"avatar"`
	Avatarmedium    interface{} `json:"avatarmedium"`
	Avatarfull      interface{} `json:"avatarfull"`
	Profileurl      interface{} `json:"profileurl"`
	Personaname     interface{} `json:"personaname"`
	LastLogin       interface{} `json:"last_login"`
	FullHistoryTime interface{} `json:"full_history_time"`
	Cheese          interface{} `json:"cheese"`
	FhUnavailable   interface{} `json:"fh_unavailable"`
	Loccountrycode  interface{} `json:"loccountrycode"`
	LastMatchTime   interface{} `json:"last_match_time"`
	Name            string      `json:"name"`
	CountryCode     string      `json:"country_code"`
	FantasyRole     int         `json:"fantasy_role"`
	TeamID          int         `json:"team_id"`
	TeamName        interface{} `json:"team_name"`
	TeamTag         interface{} `json:"team_tag"`
	IsLocked        bool        `json:"is_locked"`
	IsPro           bool        `json:"is_pro"`
	LockedUntil     int         `json:"locked_until"`
}


