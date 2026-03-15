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
					switch s.Factions[attacker].Type {
					case "enterprise":
						if s.Factions[defender].Resources.Credits > 5 {
							s.Factions[defender].War = true
							fmt.Println(s.Factions[defender].Name, "vs", s.Factions[attacker].Name)

							for s.Factions[attacker].War == true && s.Factions[defender].War == true {
								a, b := steal(s.Factions[attacker].Resources.Credits, s.Factions[defender].Resources.Credits)
								if b == 0 {
									s.Factions[attacker].War = false
									s.Factions[defender].War = false
									endwar(s, attacker)
								}
								s.Factions[attacker].Resources.Credits = a
								s.Factions[defender].Resources.Credits = b
								fmt.Println(a, b)

							}
						}
					case "collectif":
						if s.Factions[defender].Resources.Data > 5 {
							s.Factions[defender].War = true
							fmt.Println(s.Factions[defender].Name, "vs", s.Factions[attacker].Name)

							for s.Factions[attacker].War == true && s.Factions[defender].War == true {
								a, b := steal(s.Factions[attacker].Resources.Data, s.Factions[defender].Resources.Data)
								if b == 0 {
									s.Factions[attacker].War = false
									s.Factions[defender].War = false
									endwar(s, attacker)
								}
								s.Factions[attacker].Resources.Data = a
								s.Factions[defender].Resources.Data = b
								fmt.Println(a, b)

							}
						}
					case "mafia":
						if s.Factions[defender].Resources.Credits > 5 {
							s.Factions[defender].War = true
							fmt.Println(s.Factions[defender].Name, "vs", s.Factions[attacker].Name)

							for s.Factions[attacker].War == true && s.Factions[defender].War == true {
								a, b := steal(s.Factions[attacker].Resources.Credits, s.Factions[defender].Resources.Credits)
								if b == 0 {
									s.Factions[attacker].War = false
									s.Factions[defender].War = false
									endwar(s, attacker)
								}
								s.Factions[attacker].Resources.Credits = a
								s.Factions[defender].Resources.Credits = b
								fmt.Println(a, b)

							}
						}
					case "classified":
						if s.Factions[defender].Resources.Credits > 5 {
							s.Factions[defender].War = true
							fmt.Println(s.Factions[defender].Name, "vs", s.Factions[attacker].Name)

							for s.Factions[attacker].War == true && s.Factions[defender].War == true {
								a, b := steal(s.Factions[attacker].Resources.Credits, s.Factions[defender].Resources.Credits)
								if b == 0 {
									s.Factions[attacker].War = false
									s.Factions[defender].War = false
									endwar(s, attacker)
								}
								s.Factions[attacker].Resources.Credits = a
								s.Factions[defender].Resources.Credits = b
								fmt.Println(a, b)

							}
						}
					}
				}
			}
			s.Factions[attacker].War = false
		}
	}
}

func steal(attRes int, defRes int) (int, int) { // Steal resource during ATTACK
	amount := defRes / 10

	if amount == 0 {
		amount = 1
	}

	newDefRes := defRes - amount
	newAttRes := attRes + amount
	return newAttRes, newDefRes
}
