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
	Name       string
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
	Resources  Resources
	Members    int    //Nombre d'adhérents
	Type       string //Entreprise / collectif / mafia
	War        bool
	Memory     []Memory
	Alive      bool
}

type Resources struct {
	Data      int
	Influence int
	Credits   int
}

type Memory struct {
	Age   int
	Where string
	Who   string
	What  string
}
