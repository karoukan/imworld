package main

import (
	"fmt"
	"math/rand"
)

func endwar(s *Sector, i int) {
	s.Factions[i].Resources.Credits = 2
	s.Factions[i].Strength = 2
}

func war(s *Sector, attacker int) {
	fight := rand.Intn(20)

	if fight > 17 {
		if s.Factions[attacker].War != true {
			s.Factions[attacker].War = true
			s.Harvest = false
			fmt.Println(s.Factions[attacker].Name, " PREPARE FOR WAR")

			for defender := 0; defender < len(s.Factions); defender++ {
				if attacker != defender && s.Factions[defender].War == false {
					if s.Factions[attacker].Type == "enterprise" {
						if s.Factions[defender].Resources.Credits > 5 {
							s.Factions[defender].War = true
							fmt.Println(s.Factions[defender].Name, "vs", s.Factions[attacker].Name)

							for s.Factions[attacker].War == true && s.Factions[defender].War == true {
								resourcesSteal := 0
								for s.Factions[defender].Resources.Credits >= 5 {
									// fmt.Println("RESOURCES(", s.Factions[i].Name, ") BEFORE STEAL", s.Factions[i].Resources)
									// fmt.Println("RESOURCES(", s.Factions[defender].Name, ") BEFORE STEAL", s.Factions[defender].Resources)
									resourcesSteal = s.Factions[defender].Resources.Credits / 10

									if resourcesSteal == 0 {
										resourcesSteal = 1
										s.Factions[defender].War = false
										s.Factions[attacker].War = false
										endwar(s, attacker)
									}

									s.Factions[defender].Resources.Credits -= resourcesSteal
									s.Factions[attacker].Resources.Credits += resourcesSteal
									// fmt.Println("RESOURCES(", s.Factions[defender].Name, ") AFTER STEAL", s.Factions[defender].Resources)
									// fmt.Println("RESOURCES(", s.Factions[i].Name, ") AFTER STEAL", s.Factions[i].Resources)
								}
							}
						}
					}
				}
			}

		}
	}
}
