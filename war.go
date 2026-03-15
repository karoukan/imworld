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
								//Fight conditions
								attackerFaction, defenderFaction := steal(s.Factions[attacker].Resources.Credits, s.Factions[defender].Resources.Credits)
								if defenderFaction == 0 {
									s.Factions[attacker].War = false
									s.Factions[defender].War = false
									endwar(s, attacker)
								}
								s.Factions[attacker].Resources.Credits = attackerFaction
								s.Factions[defender].Resources.Credits = defenderFaction
								fmt.Println(attackerFaction, defenderFaction)

							}
						}
					case "collectif":
						if s.Factions[defender].Resources.Data > 5 {
							s.Factions[defender].War = true
							fmt.Println(s.Factions[defender].Name, "vs", s.Factions[attacker].Name)

							for s.Factions[attacker].War == true && s.Factions[defender].War == true {
								//Fight conditions
								attackerFaction, defenderFaction := steal(s.Factions[attacker].Resources.Data, s.Factions[defender].Resources.Data)
								if defenderFaction == 0 {
									s.Factions[attacker].War = false
									s.Factions[defender].War = false
									endwar(s, attacker)
								}
								s.Factions[attacker].Resources.Data = attackerFaction
								s.Factions[defender].Resources.Data = defenderFaction
								fmt.Println(attackerFaction, defenderFaction)

							}
						}
					case "mafia":
						if s.Factions[defender].Resources.Credits > 5 {
							s.Factions[defender].War = true
							fmt.Println(s.Factions[defender].Name, "vs", s.Factions[attacker].Name)

							for s.Factions[attacker].War == true && s.Factions[defender].War == true {
								//Fight conditions
								attackerFaction, defenderFaction := steal(s.Factions[attacker].Resources.Credits, s.Factions[defender].Resources.Credits)
								if defenderFaction == 0 {
									s.Factions[attacker].War = false
									s.Factions[defender].War = false
									endwar(s, attacker)
								}
								s.Factions[attacker].Resources.Credits = attackerFaction
								s.Factions[defender].Resources.Credits = defenderFaction
								fmt.Println(attackerFaction, defenderFaction)

							}
						}
					case "classified":
						if s.Factions[defender].Resources.Credits > 5 && s.Factions[defender].Members > 50 {
							s.Factions[defender].War = true
							fmt.Println(s.Factions[defender].Name, "vs", s.Factions[attacker].Name)

							for s.Factions[attacker].War == true && s.Factions[defender].War == true {
								//Fight conditions
								attackerFaction, defenderFaction := steal(s.Factions[attacker].Resources.Credits, s.Factions[defender].Resources.Credits)
								attackerMembersTotal, defenderMembersTotal := memberDie(s.Factions[attacker].Members, s.Factions[defender].Members)

								if defenderFaction == 0 {
									s.Factions[attacker].War = false
									s.Factions[defender].War = false
									endwar(s, attacker)
								}

								if defenderMembersTotal <= 0 || attackerMembersTotal <= 0 {
									s.Factions[attacker].War = false
									s.Factions[defender].War = false
									endwar(s, attacker)
								}

								s.Factions[attacker].Members = attackerMembersTotal
								s.Factions[defender].Members = defenderMembersTotal

								s.Factions[attacker].Resources.Credits = attackerFaction
								s.Factions[defender].Resources.Credits = defenderFaction

								fmt.Println("TOTAL RESOURCES:", attackerFaction, defenderFaction)
								fmt.Println("DIE TOTAL:", attackerMembersTotal, defenderMembersTotal)

							}
						}
					}
				}
			}
			s.Factions[attacker].War = false
		}
	}
}

func memberDie(attMembersTotal int, defMembersTotal int) (int, int) {
	memberDieAtt := rand.Intn(10)
	memberDieDef := rand.Intn(10)

	dieAttMembers := attMembersTotal - memberDieAtt
	dieDefMembers := defMembersTotal - memberDieDef

	if dieAttMembers <= 0 {
		dieAttMembers = 1
	}

	if dieDefMembers <= 0 {
		dieDefMembers = 1
	}

	return dieAttMembers, dieDefMembers
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
