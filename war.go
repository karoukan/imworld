package main

import (
	"fmt"
	"math/rand"
)

func endwar(d *District, faction int) {
	d.Factions[faction].Resources.Credits = 2
	d.Factions[faction].Strength = 2
}

func state(w *World, d *District, defender int, attacker int, bonusWar int) {
	d.Factions[attacker].War = true
	fmt.Println(d.Factions[defender].Name, "vs", d.Factions[attacker].Name)
	baseBonusWar := bonusWar
	for d.Factions[attacker].War == true && d.Factions[defender].War == true {
		//Fight conditions
		bonusWar = baseBonusWar * 100 / d.Factions[defender].Members
		attackerFaction, defenderFaction := steal(d.Factions[attacker].Resources.Credits, d.Factions[defender].Resources.Credits, bonusWar)
		attackerMembersTotal, defenderMembersTotal := memberDie(w, d, attacker, defender, d.Factions[attacker].Members, d.Factions[defender].Members)

		// if defenderFaction == 0 {
		// 	d.Factions[attacker].War = false
		// 	d.Factions[defender].War = false
		// 	endwar(d, attacker)
		// 	break
		// }
		if attackerMembersTotal <= 0 {
			d.Factions[attacker].War = false
			d.Factions[defender].War = false

			endwar(d, attacker)
			break
		}

		if defenderMembersTotal <= 0 {
			d.Factions[attacker].War = false
			d.Factions[defender].War = false
			endwar(d, defender)
			break
		}

		d.Factions[attacker].Members = attackerMembersTotal
		d.Factions[defender].Members = defenderMembersTotal

		d.Factions[attacker].Resources.Credits = attackerFaction
		d.Factions[defender].Resources.Credits = defenderFaction

		if d.Factions[attacker].Resources.Credits <= 1 || d.Factions[defender].Resources.Credits <= 1 {
			d.Factions[attacker].War = false
			d.Factions[defender].War = false
			break
		}

		fmt.Println("TOTAL RESOURCES:", attackerFaction, defenderFaction)
		fmt.Println("DIE TOTAL:", attackerMembersTotal, defenderMembersTotal)

	}
}

func war(w *World, s *Sector, d *District, attacker int) {
	s.Harvest = false
	for defender := 0; defender < len(d.Factions); defender++ {
		if rand.Intn(20) > 17 {

			bonusWar := warMemory(d, d.Factions[attacker].Name, defender)

			if attacker != defender && d.Factions[defender].War == false && d.Factions[defender].Alive == true {
				fmt.Println(d.Factions[attacker].Name, " PREPARE FOR WAR")
				d.Factions[defender].War = true

				switch d.Factions[attacker].Type {
				case "enterprise":
					if d.Factions[defender].Resources.Credits > 5 {
						state(w, d, defender, attacker, bonusWar)
					}
				case "collectif":
					if d.Factions[defender].Resources.Data > 5 {
						state(w, d, defender, attacker, bonusWar)
					}
				case "mafia":
					if d.Factions[defender].Resources.Credits > 5 && d.Factions[defender].Members > 100 {
						state(w, d, defender, attacker, bonusWar)
					}
				case "classified":
					if d.Factions[defender].Resources.Credits > 5 && d.Factions[defender].Members > 50 {
						state(w, d, defender, attacker, bonusWar)
					}
				}

				d.Factions[defender].Memory = append(d.Factions[defender].Memory, Memory{
					Age:   w.WorldTimer,
					Who:   d.Factions[attacker].Name,
					Where: d.Name,
					What:  "Attack",
				})
			}
		}
	}
	d.Factions[attacker].War = false
}

func warMemory(d *District, attacker string, defender int) int {
	count := 0

	for n := 0; n < len(d.Factions[defender].Memory); n++ {
		if attacker == d.Factions[defender].Memory[n].Who {
			count++
		}
	}

	return d.Factions[defender].Strength + count
}

func memberDie(w *World, d *District, attacker, defender, attMembersTotal int, defMembersTotal int) (int, int) {
	memberDieAtt := rand.Intn(10)
	memberDieDef := rand.Intn(10)

	dieAttMembers := attMembersTotal - memberDieAtt
	dieDefMembers := defMembersTotal - memberDieDef

	if dieAttMembers <= 0 {
		d.Factions[attacker].Alive = false
		dieAttMembers = 0
		controlledByChangeName(w, d, attacker, defender)
		fmt.Println(d.Factions[attacker].Name, "HAS GONE")
	}

	if dieDefMembers <= 0 {
		d.Factions[defender].Alive = false
		dieDefMembers = 0
		controlledByChangeName(w, d, defender, attacker)
		fmt.Println(d.Factions[defender].Name, "HAS GONE")
	}

	return dieAttMembers, dieDefMembers
}

func controlledByChangeName(w *World, d *District, owner int, newOwner int) {
	for AllInfrastructure := 0; AllInfrastructure < len(d.Infrastructures); AllInfrastructure++ {
		if d.Infrastructures[AllInfrastructure].ControlledBy == d.Factions[owner].Name {
			d.Infrastructures[AllInfrastructure].ControlledBy = d.Factions[newOwner].Name
			fmt.Println("INFRA TAKEN")
			d.Factions[owner].Memory = append(d.Factions[owner].Memory, Memory{
				Age:   w.WorldTimer,
				Who:   d.Factions[newOwner].Name,
				Where: d.Name,
				What:  "Infrastructure lost",
			})
		}
	}
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
