package main

import (
	"fmt"
	"math/rand"
)

func randomevent(world *World, s int) {
	if rand.Intn(100) <= 10 {
		ev := rand.Intn(len(world.Sectors[s].Events))

		impactFaction := world.Sectors[s].Events[ev].Impact
		world.Sectors[s].Population -= impactFaction * 3

		fmt.Println("POP:", world.Sectors[s].Events[ev].Name, "(", world.Sectors[s].Population, ")")
	}
}

func population(world *World, s int) {
	world.Sectors[s].Population += world.Sectors[s].Population / 100
	fmt.Println("POP INCREASE:", world.Sectors[s].Population)
}
