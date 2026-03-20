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

func population(d *District) int {
	increaseChange := 40

	miseryApply := misery(d)

	if miseryApply > 25 {
		increaseChange -= 30
	}

	if rand.Intn(100) <= increaseChange {
		pop := d.Population + d.Population/100

		if increaseChange == 30 {
			fmt.Println("POP(", d.Name, ") MALUS :", d.Population)
		}
		fmt.Println("POP(", d.Name, ") INCREASE:", d.Population)

		return pop
	}

	return d.Population
}

func populationJoinFaction(world *World, sector int, district int, faction int, pop int) {
	newMembers := rand.Intn(pop) / 100

	if newMembers > 0 {
		world.Sectors[sector].Districts[district].Factions[faction].Members += newMembers
	}
}

func misery(d *District) int {
	creditPerFaction := 0
	for allCreditsFaction := 0; allCreditsFaction < len(d.Factions); allCreditsFaction++ {
		creditPerFaction += d.Factions[allCreditsFaction].Resources.Credits

	}
	if len(d.Factions) == 0 {
		d.Misery = 1
		return d.Misery
	}

	totalCreditPerFaction := creditPerFaction / len(d.Factions)

	if totalCreditPerFaction == 0 {
		totalCreditPerFaction = 1
	}

	// miseryIndex := 100 - (totalCreditPerFaction * 100 / d.Population)
	d.Misery = max(100-totalCreditPerFaction, 0)
	fmt.Println("MISERY INDEX", d.Misery)
	return d.Misery
}
