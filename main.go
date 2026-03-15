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
				Size:       10,
				Population: 10000,
				Location:   1,
				Factions: []Faction{
					{
						id:         1,
						Name:       "MegaCorp",
						Strength:   2,
						Ideology:   "Neutral",
						Reputation: 1,
						Type:       "enterprise",
						Members:    100,
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
						Resources: Resources{
							Credits:   10,
							Influence: 4,
							Data:      12,
						},
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
				Factions: []Faction{
					{
						id:         1,
						Name:       "Nova",
						Strength:   8,
						Ideology:   "Loyal",
						Reputation: 1,
						Members:    100,
						Type:       "classified",
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
						Resources: Resources{
							Credits:   10,
							Influence: 10,
							Data:      2,
						},
					},
				},
			},
		},
	}

	fmt.Println(world)

	for {
		world.WorldTimer++
		time.Sleep(time.Second)

		for s := 0; s < len(world.Sectors); s++ {
			randomevent(&world, s)

			newPopulation := population(&world, s)
			world.Sectors[s].Population = newPopulation

			for k := 0; k < len(world.Sectors[s].Events); k++ {
				for i := 0; i < len(world.Sectors[s].Factions); i++ {

					membersJoined := rand.Intn(100)
					if membersJoined > 17 && newPopulation >= 100 {
						populationJoinFaction(&world, s, i, newPopulation)
						fmt.Println(world.Sectors[s].Factions[i].Name, "TOTAL MEMBER:", world.Sectors[s].Factions[i].Members)
					}

					if world.Sectors[s].Harvest == true {
						gathering(&world, s, i)
						fmt.Println(world.Sectors[s].Factions[i].Name, "GATHERING:", world.Sectors[s].Factions[i].Resources)
					}

					// if world.Sectors[s].Factions[i].Strength >= 2 && world.Sectors[s].Harvest == true {
					// 	//world.Sectors[s].Factions[i].Resources += +world.Sectors[s].Factions[i].Strength
					// 	// fmt.Println(world.Sectors[s].Factions[i].Name, "GATHERING:", world.Sectors[s].Factions[i].Resources)
					// 	for
					// }
					war(&world.Sectors[s], i)
					world.Sectors[s].Harvest = true
				}
			}
		}
	}
}
