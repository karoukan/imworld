package main

type World struct {
	WorldTimer int
	Sectors    []Sector
}

type Event struct {
	Name   string
	Type   string
	Impact int
}

type Sector struct {
	Harvest    bool
	Fight      bool
	Size       int
	Population int
	Location   int //peut etre mettre un float si on veut des coord GPS à voir plus tard ?
	Factions   []Faction
	Events     []Event
}

type Faction struct {
	id         int
	Name       string
	Strength   int
	Ideology   string
	Reputation int
	Resources  int
	War        bool
}
