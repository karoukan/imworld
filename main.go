package main

import (
	"fmt"
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
						Resources:  3,
					},
					{
						id:         2,
						Name:       "GHOST",
						Strength:   3,
						Ideology:   "Xeno",
						Reputation: 1,
						Resources:  5,
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
						Resources:  10,
					},
					{
						id:         2,
						Name:       "EmpirzSec",
						Strength:   6,
						Ideology:   "Bad",
						Reputation: 1,
						Resources:  5,
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
			population(&world, s)
			for k := 0; k < len(world.Sectors[s].Events); k++ {
				for i := 0; i < len(world.Sectors[s].Factions); i++ {
					if world.Sectors[s].Factions[i].Strength >= 2 && world.Sectors[s].Harvest == true {
						world.Sectors[s].Factions[i].Resources += +world.Sectors[s].Factions[i].Strength
						// fmt.Println(world.Sectors[s].Factions[i].Name, "GATHERING:", world.Sectors[s].Factions[i].Resources)
					}
					war(&world.Sectors[s], i)
					world.Sectors[s].Harvest = true
				}
			}
		}
	}
}
