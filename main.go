package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Imworld")
	world := World{
		Sectors: []Sector{
			{
				Harvest: true,
				Fight:   false,
				Events: []Event{
					{
						Name:   "Pandemia",
						Type:   "Virus",
						Impact: 3,
					},
					{
						Name:   "Earh quake",
						Type:   "Catastrophe",
						Impact: 4,
					},
				},
				Name:       "Sector A",
				Size:       2,
				Population: 10000,
				Location:   1,
				Districts: []District{
					{
						Name:       "Ker-Uhel",
						Population: 5000,
						Size:       1,
						Infrastructures: []Infrastructure{{
							Name:  "Monee",
							Type:  "Bank",
							State: "Ready",
							InUse: true,
						}},
						Factions: []Faction{
							{
								id:         1,
								Name:       "MegaCorp",
								Strength:   2,
								Ideology:   "Neutral",
								Reputation: 1,
								Type:       "enterprise",
								Members:    100,
								Alive:      true,
								Memory: []Memory{
									{
										Age:   4,
										Where: "Sector A",
										Who:   "GHOST",
										What:  "Attack",
									},
								},
								Resources: Resources{
									Credits:   10,
									Influence: 15,
									Data:      2,
								},
							},
							{
								id:         2,
								Name:       "GHOST",
								Strength:   3,
								Ideology:   "Xeno",
								Reputation: 1,
								Members:    100,
								Type:       "collectif",
								Alive:      true,
								Resources: Resources{
									Credits:   10,
									Influence: 4,
									Data:      12,
								},
							},
						},
					},
					{
						Name:       "Ar Santé",
						Population: 5000,
						Size:       1,
						Infrastructures: []Infrastructure{{
							Name:  "L'entrepot",
							Type:  "Warehouse",
							State: "Ready",
							InUse: true,
						}},
						Factions: []Faction{},
					},
				},
			},
			{
				Harvest: true,
				Fight:   false,
				Events: []Event{
					{
						Name:   "Pandemia",
						Type:   "Virus",
						Impact: 3,
					},
					{
						Name:   "Earh quake",
						Type:   "Catastrophe",
						Impact: 4,
					},
				},
				Name:       "CITY 17",
				Size:       5,
				Population: 1000,
				Location:   2,
				Districts: []District{
					{
						Name:       "White Mesa",
						Population: 450,
						Size:       2,
						Infrastructures: []Infrastructure{{
							Name:  "Anormal Material Labs",
							Type:  "Research Center",
							State: "Maintenance",
							InUse: true,
						}},
						Factions: []Faction{
							{
								id:         1,
								Name:       "Nova",
								Strength:   8,
								Ideology:   "Loyal",
								Reputation: 1,
								Members:    100,
								Type:       "classified",
								Alive:      true,
								Resources: Resources{
									Credits:   10,
									Influence: 4,
									Data:      12,
								},
							},
							{
								id:         2,
								Name:       "EmpirzSec",
								Strength:   6,
								Ideology:   "Bad",
								Reputation: 1,
								Type:       "mafia",
								Members:    100,
								Alive:      true,
								Resources: Resources{
									Credits:   10,
									Influence: 10,
									Data:      2,
								},
							},
						},
					},
					{
						Name:       "Nex",
						Population: 550,
						Size:       3,
						Infrastructures: []Infrastructure{{
							Name:  "Freewoman",
							Type:  "Factory",
							State: "Building",
							InUse: false,
						}},
						Factions: []Faction{},
					},
				},
			},
		},
	}

	fmt.Println(world)

	for {
		world.WorldTimer++
		time.Sleep(time.Second)

		for AllSectors := 0; AllSectors < len(world.Sectors); AllSectors++ { //Fetch sur tous les secteurs
			randomevent(&world, AllSectors)

			newPopulation := population(&world, AllSectors)
			world.Sectors[AllSectors].Population = newPopulation

			for AllDistricts := 0; AllDistricts < len(world.Sectors[AllSectors].Districts); AllDistricts++ { //Fetch sur tous les districts
				district_l := &world.Sectors[AllSectors].Districts[AllDistricts]
				// sector_l := &world.Sectors[AllSectors]
				for AllFactions := 0; AllFactions < len(district_l.Factions); AllFactions++ { //Fetch sur toutes les factions
					membersJoined := rand.Intn(100)
					if membersJoined > 17 && newPopulation >= 100 && district_l.Factions[AllFactions].Alive == true {
						populationJoinFaction(&world, AllSectors, AllDistricts, AllFactions, newPopulation)
						fmt.Println(district_l.Factions[AllFactions].Name, "TOTAL MEMBER:", district_l.Factions[AllFactions].Members)
					}

					if world.Sectors[AllSectors].Harvest == true && district_l.Factions[AllFactions].Alive == true {
						gathering(&world, AllSectors, AllDistricts, AllFactions)
						fmt.Println(district_l.Factions[AllFactions].Name, "GATHERING:", district_l.Factions[AllFactions].Resources)
					}

					// if world.Sectors[s].Factions[i].Strength >= 2 && world.Sectors[s].Harvest == true {
					// 	//world.Sectors[s].Factions[i].Resources += +world.Sectors[s].Factions[i].Strength
					// 	// fmt.Println(world.Sectors[s].Factions[i].Name, "GATHERING:", world.Sectors[s].Factions[i].Resources)
					// 	for
					// }

					if district_l.Factions[AllFactions].Alive == true {
						// initTrade(&world, &world.Sectors[AllSectors].Districts[AllDistricts], AllFactions)
						// war(&world, &world.Sectors[AllSectors], &world.Sectors[AllSectors].Districts[AllDistricts], AllFactions)
						action := decide(&world, &world.Sectors[AllFactions], district_l, AllFactions)
						fmt.Println(action)
					}
				}
				world.Sectors[AllSectors].Harvest = true
			}
		}
	}
}
