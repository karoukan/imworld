package main

import (
	"fmt"
	"math/rand"
)

func endwar(s *Sector, i int) {
	s.Factions[i].Resources = 2
	s.Factions[i].Strength = 2
}

func war(s *Sector, i int) {
	fight := rand.Intn(20)

	if fight > 17 {
		if s.Factions[i].Resources >= 20 && s.Factions[i].War != true {
			s.Factions[i].War = true
			s.Harvest = false
			fmt.Println(s.Factions[i].Name, ": WORLD DOMINATION :", s.Factions[i].War)

			for j := 0; j < len(s.Factions); j++ {
				if i != j && s.Factions[j].War == false && s.Factions[j].Resources > 5 {
					s.Factions[j].War = true
					fmt.Println(s.Factions[j].Name, "vs", s.Factions[i].Name)

					for s.Factions[i].War == true && s.Factions[j].War == true {
						resourcesSteal := 0
						for s.Factions[j].Resources >= 5 {
							// fmt.Println("RESOURCES(", s.Factions[i].Name, ") BEFORE STEAL", s.Factions[i].Resources)
							// fmt.Println("RESOURCES(", s.Factions[j].Name, ") BEFORE STEAL", s.Factions[j].Resources)
							resourcesSteal = s.Factions[j].Resources / 10

							if resourcesSteal == 0 {
								resourcesSteal = 1
								s.Factions[j].War = false
								s.Factions[i].War = false
								endwar(s, i)
							}

							s.Factions[j].Resources -= resourcesSteal
							s.Factions[i].Resources += resourcesSteal
							// fmt.Println("RESOURCES(", s.Factions[j].Name, ") AFTER STEAL", s.Factions[j].Resources)
							// fmt.Println("RESOURCES(", s.Factions[i].Name, ") AFTER STEAL", s.Factions[i].Resources)
						}
					}
				}
			}
		}
	}
}
