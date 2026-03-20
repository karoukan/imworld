package main

type World struct {
	WorldTimer int
	Sectors    []Sector
	Government Gov
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
	Districts  []District
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

type Gov struct {
	Name      string
	Resources Resources
	Members   int //Nombre d'adhérents
	Taxe      int
	Memory    []Memory
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

type District struct {
	Population      int
	Name            string
	Size            int
	Location        int
	Misery          int
	Factions        []Faction
	Infrastructures []Infrastructure
}

type Infrastructure struct {
	Name string
	Type string
	// GeneratedRessources string
	State        string //Maintenance, Ready, Building (+ de states plus tard)
	InUse        bool   //Signifie que des gens y travaille (Ready ou Maintenance), Building (en cours de construction donc false)
	ControlledBy string
}
