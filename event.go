package main

import (
	"fmt"
	"math/rand"
)

func randomevent(world *World, s int) {
	if rand.Intn(100) <= 1 {
		ev := rand.Intn(len(world.Sectors[s].Events))
		impactFaction := world.Sectors[s].Events[ev].Impact

		if world.Sectors[s].Events[ev].Type == "Virus" {
			world.Sectors[s].Population -= impactFaction * 3
			fmt.Println("POP:", world.Sectors[s].Events[ev].Name, "(", world.Sectors[s].Population, ")")
		}

		if world.Sectors[s].Events[ev].Type == "Catastrophe" {
			for k := 0; k < len(world.Sectors[s].Factions); k++ {
				world.Sectors[s].Factions[k].Strength -= impactFaction - 2
				if world.Sectors[s].Factions[k].Strength <= 0 {
					world.Sectors[s].Factions[k].Strength = 1
				}
				fmt.Println("STRENGTH (", world.Sectors[s].Factions[k].Name, ") DECREASE:", world.Sectors[s].Factions[k].Strength)
			}
		}
	}
}

func population(world *World, s int) int {
	if rand.Intn(100) <= 40 {
		pop := world.Sectors[s].Population + world.Sectors[s].Population/100
		fmt.Println("POP(", world.Sectors[s].Name, ") INCREASE:", world.Sectors[s].Population)

		return pop
	}

	return world.Sectors[s].Population
}

func populationJoinFaction(world *World, sector int, district int, faction int, pop int) {
	newMembers := rand.Intn(pop) / 100

	if newMembers > 0 {
		world.Sectors[sector].Districts[district].Factions[faction].Members += newMembers
	}
}
