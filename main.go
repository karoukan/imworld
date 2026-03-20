package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Imworld")
	world := World{
		Government: Gov{
			Name: "ARCH", //Authority for Regional Civic Harmony
			Resources: Resources{
				Credits:   0,
				Influence: 0,
				Data:      0,
			},
			Members: 100,
			Taxe:    2,
		},
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
				Name:       "Sector A",
				Size:       2,
				Population: 10000,
				Location:   1,
				Districts: []District{
					{
						Name:       "Ker-Uhel",
						Population: 5000,
						Size:       1,
						Infrastructures: []Infrastructure{
							{
								Name:         "Monee",
								Type:         "Bank",
								State:        "Ready",
								InUse:        true,
								ControlledBy: "MegaCorp",
							},
							{
								Name:         "Le Super Jaudy",
								Type:         "Market",
								State:        "Ready",
								InUse:        true,
								ControlledBy: "MegaCorp",
							},
						},
						Factions: []Faction{
							{
								id:         1,
								Name:       "MegaCorp",
								Strength:   2,
								Ideology:   "Neutral",
								Reputation: 1,
								Type:       "enterprise",
								Members:    100,
								Alive:      true,
								Memory: []Memory{
									{
										Age:   4,
										Where: "Sector A",
										Who:   "GHOST",
										What:  "Attack",
									},
								},
								Resources: Resources{
									Credits:   10,
									Influence: 15,
									Data:      2,
								},
							},
							{
								id:         2,
								Name:       "GHOST",
								Strength:   3,
								Ideology:   "Xeno",
								Reputation: 1,
								Members:    100,
								Type:       "collectif",
								Alive:      true,
								Resources: Resources{
									Credits:   10,
									Influence: 4,
									Data:      12,
								},
							},
						},
					},
					{
						Name:       "Ar Santé",
						Population: 5000,
						Size:       1,
						Infrastructures: []Infrastructure{{
							Name:         "Odyssea",
							Type:         "Datacenter",
							State:        "Ready",
							InUse:        true,
							ControlledBy: "",
						}},
						Factions: []Faction{},
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
				Name:       "CITY 17",
				Size:       5,
				Population: 1000,
				Location:   2,
				Districts: []District{
					{
						Name:       "White Mesa",
						Population: 450,
						Size:       2,
						Infrastructures: []Infrastructure{{
							Name:         "Anormal Material Labs",
							Type:         "Research Center",
							State:        "Maintenance",
							ControlledBy: "Nova",
							InUse:        true,
						}},
						Factions: []Faction{
							{
								id:         1,
								Name:       "Nova",
								Strength:   8,
								Ideology:   "Loyal",
								Reputation: 1,
								Members:    100,
								Type:       "classified",
								Alive:      true,
								Resources: Resources{
									Credits:   10,
									Influence: 4,
									Data:      12,
								},
							},
							{
								id:         2,
								Name:       "EmpirzSec",
								Strength:   6,
								Ideology:   "Bad",
								Reputation: 1,
								Type:       "mafia",
								Members:    100,
								Alive:      true,
								Resources: Resources{
									Credits:   10,
									Influence: 10,
									Data:      2,
								},
							},
						},
					},
					{
						Name:       "Nex",
						Population: 550,
						Size:       3,
						Infrastructures: []Infrastructure{{
							Name:         "Freewoman",
							Type:         "Factory",
							State:        "Building",
							InUse:        false,
							ControlledBy: "",
						}},
						Factions: []Faction{},
					},
				},
			},
		},
	}

	fmt.Println(world)

	for {
		world.WorldTimer++
		time.Sleep(time.Second)

		for AllSectors := 0; AllSectors < len(world.Sectors); AllSectors++ { //Fetch sur tous les secteurs
			randomevent(&world, AllSectors)

			for AllDistricts := 0; AllDistricts < len(world.Sectors[AllSectors].Districts); AllDistricts++ { //Fetch sur tous les districts
				district_l := &world.Sectors[AllSectors].Districts[AllDistricts]
				newPopulation := population(district_l)
				district_l.Population = newPopulation

				// si un district n'a pas faction alors on peut ajouter une nouvelle faction
				aliveCount := 0
				for i := 0; i < len(district_l.Factions); i++ {
					if district_l.Factions[i].Alive == true {
						aliveCount++
					}
				}

				if len(district_l.Factions) == 0 || aliveCount < 2 {

					names := []string{"Nexus", "Vortex", "Syndicate", "Phantom", "Onyx"}
					types := []string{"enterprise", "collectif", "mafia", "classified"}
					randomName := rand.Intn(len(names))
					randomType := rand.Intn(len(types))
					nameAlreadyTaken := false
					for checkFactionName := 0; checkFactionName < len(district_l.Factions); checkFactionName++ {
						if district_l.Factions[checkFactionName].Name == names[randomName] {
							nameAlreadyTaken = true
							break
						}
					}
					if nameAlreadyTaken == false {
						district_l.Factions = append(district_l.Factions, Faction{
							Name:       names[randomName],
							Strength:   rand.Intn(6),
							Ideology:   "Bad",
							Reputation: 1,
							Type:       types[randomType],
							Members:    rand.Intn(50),
							Alive:      true,
							Resources: Resources{
								Credits:   rand.Intn(10),
								Influence: rand.Intn(10),
								Data:      rand.Intn(10),
							},
						})
					}

				}

				for AllInfrastructure := 0; AllInfrastructure < len(district_l.Infrastructures); AllInfrastructure++ {
					if len(district_l.Infrastructures[AllInfrastructure].ControlledBy) != 0 {
						for nFaction := 0; nFaction < len(district_l.Factions); nFaction++ {

							if district_l.Factions[nFaction].Name == district_l.Infrastructures[AllInfrastructure].ControlledBy {
								if district_l.Infrastructures[AllInfrastructure].InUse == true {
									switch district_l.Infrastructures[AllInfrastructure].Type {
									case "Bank":
										district_l.Factions[nFaction].Resources.Credits += 10
										fmt.Println(district_l.Infrastructures[AllInfrastructure].ControlledBy, "CONTROLED", district_l.Infrastructures[AllInfrastructure].Name, "HAS GAINED MORE CREDITS", district_l.Factions[nFaction].Resources.Credits)
									case "Factory":
										district_l.Factions[nFaction].Resources.Credits += 5
									case "Market":
										district_l.Factions[nFaction].Resources.Credits += 5
									case "Datacenter":
										district_l.Factions[nFaction].Resources.Data += 10
									case "Research Center":
										if district_l.Factions[nFaction].Resources.Credits <= 4 {
											district_l.Infrastructures[AllInfrastructure].InUse = false
											break
										}

										district_l.Factions[nFaction].Resources.Data += 15
										district_l.Factions[nFaction].Resources.Credits -= 5
									}
								}

								afterfee, feePocketGov := tax(district_l, &world, nFaction)
								district_l.Factions[nFaction].Resources.Credits = afterfee
								world.Government.Resources.Credits = feePocketGov
								fmt.Println(world.Government.Name, "GOV TAXE ", afterfee, feePocketGov)
							}
						}
					}
				}
				for AllFactions := 0; AllFactions < len(district_l.Factions); AllFactions++ { //Fetch sur toutes les factions
					membersJoined := rand.Intn(100)

					if membersJoined > 17 && newPopulation >= 100 && district_l.Factions[AllFactions].Alive == true {
						populationJoinFaction(&world, AllSectors, AllDistricts, AllFactions, newPopulation)
						fmt.Println(district_l.Factions[AllFactions].Name, "TOTAL MEMBER:", district_l.Factions[AllFactions].Members)
					}

					if world.Sectors[AllSectors].Harvest == true && district_l.Factions[AllFactions].Alive == true {
						gathering(&world, AllSectors, AllDistricts, AllFactions)
						fmt.Println(district_l.Factions[AllFactions].Name, "GATHERING:", district_l.Factions[AllFactions].Resources)
					}

					misery(district_l)

					// if world.Sectors[s].Factions[i].Strength >= 2 && world.Sectors[s].Harvest == true {
					// 	//world.Sectors[s].Factions[i].Resources += +world.Sectors[s].Factions[i].Strength
					// 	// fmt.Println(world.Sectors[s].Factions[i].Name, "GATHERING:", world.Sectors[s].Factions[i].Resources)
					// 	for
					// }

					if district_l.Factions[AllFactions].Alive == true {
						for AllInfrastructure := 0; AllInfrastructure < len(district_l.Infrastructures); AllInfrastructure++ { //Fetch toutes les infra
							//On fetch toutes les infra pour filter celle qui sont dans mon district et qui n'ont pas de propriétaire
							//Si on en trouve on se l'approprie si on peut payer 50 cr sinon on passe

							if len(district_l.Infrastructures[AllInfrastructure].ControlledBy) == 0 { // !=0 pour avoir les infra controlled et ==0 pour les dispo
								if district_l.Factions[AllFactions].Resources.Credits >= 51 {
									district_l.Factions[AllFactions].Resources.Credits -= 50
									district_l.Infrastructures[AllInfrastructure].ControlledBy = district_l.Factions[AllFactions].Name
									fmt.Println(district_l.Infrastructures[AllInfrastructure].ControlledBy, "ACQUIRED ", district_l.Infrastructures[AllInfrastructure].Name)
								}
							}

						}

						action := decide(&world, &world.Sectors[AllSectors], district_l, AllFactions)
						fmt.Println(action)
					}
				}
			}
			world.Sectors[AllSectors].Harvest = true
		}
	}
}
