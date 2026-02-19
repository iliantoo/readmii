package main

// Models CS2
type ActionDetails struct {
	BindRequired bool   `json:"bind_required"`
	BindCommand  string `json:"bind_command,omitempty"`
	Movement     string `json:"movement"`
	ClickType    string `json:"click_type"`
}

type Media struct {
	PositionImage      string `json:"position_image"`
	PositionThumbnail  string `json:"position_thumbnail"`
	CrosshairImage     string `json:"crosshair_image"`
	CrosshairZoomLevel string `json:"crosshair_zoom_level"`
	DemoGif            string `json:"demo_gif"`
	DemoThumbnail      string `json:"demo_thumbnail"`
}

type Lineup struct {
	ID            string        `json:"id"`
	Title         string        `json:"title"`
	Description   string        `json:"description"`
	GrenadeType   string        `json:"grenade_type"`
	Side          string        `json:"side"`
	Difficulty    string        `json:"difficulty"`
	ThrowZone     string        `json:"throw_zone"`
	LandingZone   string        `json:"landing_zone"`
	ActionType    string        `json:"action_type"`
	ActionDetails ActionDetails `json:"action_details"`
	Media         Media         `json:"media"`
	Tags          []string      `json:"tags"`
	Popularity    float64       `json:"popularity"`
	Views         int           `json:"views"`
}

type Map struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	DisplayName  string   `json:"display_name"`
	CoverImage   string   `json:"cover_image"`
	TotalLineups int      `json:"total_lineups"`
	Lineups      []Lineup `json:"lineups"`
}

type Library struct {
	Maps []Map `json:"maps"`
}
