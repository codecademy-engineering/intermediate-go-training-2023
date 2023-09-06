package models

type furColor struct {
	Primary    string
	Highlights string
}

type coordinates struct {
	Lat float64
	Lng float64
}

type Squirrel struct {
	ID                     string   `json:"ID"`
	FurColor               furColor `json:"Fur_Color"`
	Park                   string
	InteractionsWithHumans string `json:"Interactions_With_Humans"`
	Activities             []string
	Coorditantes           coordinates
}
