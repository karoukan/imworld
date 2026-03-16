package main

import (
	"fmt"
	"math/rand"
)

func endwar(s *Sector, i int) {
	s.Factions[i].Resources.Credits = 2
	s.Factions[i].Strength = 2
}

func war(w *World, s *Sector, attacker int) {
	fight := rand.Intn(20)

	if fight > 17 {
		if s.Factions[attacker].War != true {
			s.Factions[attacker].War = true
			s.Harvest = false
			fmt.Println(s.Factions[attacker].Name, " PREPARE FOR WAR")

			for defender := 0; defender < len(s.Factions); defender++ {
				bonusWar := warMemory(s, s.Factions[attacker].Name, defender)

				if attacker != defender && s.Factions[defender].War == false && s.Factions[defender].Alive == true {
					switch s.Factions[attacker].Type {
					case "enterprise":
						if s.Factions[defender].Resources.Credits > 5 {
							s.Factions[defender].War = true
							fmt.Println(s.Factions[defender].Name, "vs", s.Factions[attacker].Name)

							for s.Factions[attacker].War == true && s.Factions[defender].War == true {
								//Fight conditions
								attackerFaction, defenderFaction := steal(s.Factions[attacker].Resources.Credits, s.Factions[defender].Resources.Credits, bonusWar)
								attackerMembersTotal, defenderMembersTotal := memberDie(s, attacker, defender, s.Factions[attacker].Members, s.Factions[defender].Members)

								if defenderFaction == 0 {
									s.Factions[attacker].War = false
									s.Factions[defender].War = false
									endwar(s, attacker)
									break
								}

								if defenderMembersTotal <= 0 || attackerMembersTotal <= 0 {
									s.Factions[attacker].War = false
									s.Factions[defender].War = false
									endwar(s, attacker)
									break
								}

								s.Factions[attacker].Members = attackerMembersTotal
								s.Factions[defender].Members = defenderMembersTotal

								s.Factions[attacker].Resources.Credits = attackerFaction
								s.Factions[defender].Resources.Credits = defenderFaction

								fmt.Println("TOTAL RESOURCES:", attackerFaction, defenderFaction)
								fmt.Println("DIE TOTAL:", attackerMembersTotal, defenderMembersTotal)

							}
						}
					case "collectif":
						if s.Factions[defender].Resources.Data > 5 {
							s.Factions[defender].War = true
							fmt.Println(s.Factions[defender].Name, "vs", s.Factions[attacker].Name)

							for s.Factions[attacker].War == true && s.Factions[defender].War == true {
								//Fight conditions
								attackerFaction, defenderFaction := steal(s.Factions[attacker].Resources.Credits, s.Factions[defender].Resources.Credits, bonusWar)
								attackerMembersTotal, defenderMembersTotal := memberDie(s, attacker, defender, s.Factions[attacker].Members, s.Factions[defender].Members)

								if defenderFaction == 0 {
									s.Factions[attacker].War = false
									s.Factions[defender].War = false
									endwar(s, attacker)
									break
								}

								if defenderMembersTotal <= 0 || attackerMembersTotal <= 0 {
									s.Factions[attacker].War = false
									s.Factions[defender].War = false
									endwar(s, attacker)
									break
								}

								s.Factions[attacker].Members = attackerMembersTotal
								s.Factions[defender].Members = defenderMembersTotal

								s.Factions[attacker].Resources.Credits = attackerFaction
								s.Factions[defender].Resources.Credits = defenderFaction

								fmt.Println("TOTAL RESOURCES:", attackerFaction, defenderFaction)
								fmt.Println("DIE TOTAL:", attackerMembersTotal, defenderMembersTotal)

							}
						}
					case "mafia":
						if s.Factions[defender].Resources.Credits > 5 && s.Factions[defender].Members > 100 {
							s.Factions[defender].War = true
							fmt.Println(s.Factions[defender].Name, "vs", s.Factions[attacker].Name)

							for s.Factions[attacker].War == true && s.Factions[defender].War == true {
								//Fight conditions
								attackerFaction, defenderFaction := steal(s.Factions[attacker].Resources.Credits, s.Factions[defender].Resources.Credits, bonusWar)
								attackerMembersTotal, defenderMembersTotal := memberDie(s, attacker, defender, s.Factions[attacker].Members, s.Factions[defender].Members)

								if defenderFaction == 0 {
									s.Factions[attacker].War = false
									s.Factions[defender].War = false
									endwar(s, attacker)
									break
								}

								if defenderMembersTotal <= 0 || attackerMembersTotal <= 0 {
									s.Factions[attacker].War = false
									s.Factions[defender].War = false
									endwar(s, attacker)
									break
								}

								s.Factions[attacker].Members = attackerMembersTotal
								s.Factions[defender].Members = defenderMembersTotal

								s.Factions[attacker].Resources.Credits = attackerFaction
								s.Factions[defender].Resources.Credits = defenderFaction

								fmt.Println("TOTAL RESOURCES:", attackerFaction, defenderFaction)
								fmt.Println("DIE TOTAL:", attackerMembersTotal, defenderMembersTotal)

							}
						}
					case "classified":
						if s.Factions[defender].Resources.Credits > 5 && s.Factions[defender].Members > 50 {
							s.Factions[defender].War = true
							fmt.Println(s.Factions[defender].Name, "vs", s.Factions[attacker].Name)

							for s.Factions[attacker].War == true && s.Factions[defender].War == true {
								//Fight conditions
								attackerFaction, defenderFaction := steal(s.Factions[attacker].Resources.Credits, s.Factions[defender].Resources.Credits, bonusWar)
								attackerMembersTotal, defenderMembersTotal := memberDie(s, attacker, defender, s.Factions[attacker].Members, s.Factions[defender].Members)

								if defenderFaction == 0 {
									s.Factions[attacker].War = false
									s.Factions[defender].War = false
									endwar(s, attacker)
									break
								}

								if defenderMembersTotal <= 0 || attackerMembersTotal <= 0 {
									s.Factions[attacker].War = false
									s.Factions[defender].War = false
									endwar(s, attacker)
									break
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

					s.Factions[defender].Memory = append(s.Factions[defender].Memory, Memory{
						Age:   w.WorldTimer,
						Who:   s.Factions[attacker].Name,
						Where: s.Name,
						What:  "Attack",
					})
				}
			}
			s.Factions[attacker].War = false
		}
	}
}

func warMemory(s *Sector, attacker string, defender int) int {
	count := 0

	for n := 0; n < len(s.Factions[defender].Memory); n++ {
		if attacker == s.Factions[defender].Memory[n].Who {
			count++
		}
	}

	return s.Factions[defender].Strength + count
}

func memberDie(s *Sector, attacker, defender, attMembersTotal int, defMembersTotal int) (int, int) {
	memberDieAtt := rand.Intn(10)
	memberDieDef := rand.Intn(10)

	dieAttMembers := attMembersTotal - memberDieAtt
	dieDefMembers := defMembersTotal - memberDieDef

	if dieAttMembers <= 0 {
		s.Factions[attacker].Alive = false
		dieAttMembers = 0
		fmt.Println(s.Factions[attacker].Name, "HAS GONE")
	}

	if dieDefMembers <= 0 {
		s.Factions[defender].Alive = false
		dieDefMembers = 0
		fmt.Println(s.Factions[defender].Name, "HAS GONE")
	}

	return dieAttMembers, dieDefMembers
}

func steal(attRes int, defRes int, bonusWar int) (int, int) { // Steal resource during ATTACK
	amount := defRes / 10

	if amount == 0 {
		amount = 1
	}

	newDefRes := defRes - amount + bonusWar

	if bonusWar >= amount {
		newDefRes = defRes - amount + bonusWar/2
	}

	newAttRes := attRes + amount
	return newAttRes, newDefRes
}
