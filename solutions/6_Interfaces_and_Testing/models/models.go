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
	ID                     string
	FurColor               furColor
	Park                   string
	InteractionsWithHumans string
	Activities             []string
	Coorditantes           coordinates
}
