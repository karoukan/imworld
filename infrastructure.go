package main

import "fmt"

func repair(d *District, w *World) {
	// Une faction répare ca propre infra sur son territoire
	// Dès qu'une infra est en maintenance, au bout de n ticks et si la faction paie (sans que ca la ruine), l'infra repasse en Ready
	// Pour le moment, l'infra ne peut pas passer en mode abandonned
	for AllInfrastructure := 0; AllInfrastructure < len(d.Infrastructures); AllInfrastructure++ {
		for faction := 0; faction < len(d.Factions); faction++ {
			if d.Infrastructures[AllInfrastructure].State == "Maintenance" {
				if d.Infrastructures[AllInfrastructure].ControlledBy == d.Factions[faction].Name {
					d.Infrastructures[AllInfrastructure].OperationsSince = w.WorldTimer

					if d.Factions[faction].Resources.Credits >= d.Infrastructures[AllInfrastructure].MaintenanceCost {

						d.Factions[faction].Resources.Credits -= d.Infrastructures[AllInfrastructure].MaintenanceCost

						d.Infrastructures[AllInfrastructure].State = "Ready"
						d.Infrastructures[AllInfrastructure].StartSince = w.WorldTimer //C'est payé donc ca leur appartient bien
						d.Infrastructures[AllInfrastructure].OperationsSince = 0       //Reset State pour éviter rachat par ARCH même si Ready lol

						fmt.Println("OPERATION MAINTENANCE COMPLETED")
					}
				}
			}
		}
	}
}

func opaHostile(d *District, w *World) {
	//Si une faction n'a pas payé par manque de moyen ALORS ARCH rachete l'infra
	for AllInfrastructure := 0; AllInfrastructure < len(d.Infrastructures); AllInfrastructure++ {
		if d.Infrastructures[AllInfrastructure].State == "Maintenance" {
			timerOpa := w.WorldTimer - d.Infrastructures[AllInfrastructure].OperationsSince
			if timerOpa > 30 {
				d.Infrastructures[AllInfrastructure].ControlledBy = w.Government.Name
				d.Infrastructures[AllInfrastructure].State = "Ready"
				d.Infrastructures[AllInfrastructure].StartSince = w.WorldTimer //C'est payé donc ca leur appartient bien
				d.Infrastructures[AllInfrastructure].OperationsSince = 0       //Reset State pour éviter rachat par ARCH même si Ready lol

				fmt.Println("ARCH ACQUIRED INFRASTRACTURE")
			}
		}
	}
}

func archMaintenance(d *District, w *World) {
	for AllInfrastructure := 0; AllInfrastructure < len(d.Infrastructures); AllInfrastructure++ {
		if w.Government.Resources.Credits <= 0 {
			if d.Infrastructures[AllInfrastructure].ControlledBy == w.Government.Name {
				d.Infrastructures[AllInfrastructure].State = "Maintenance"
				fmt.Println("ARCH LOST INFRA")
			}

		} else if d.Infrastructures[AllInfrastructure].ControlledBy == w.Government.Name {
			w.Government.Resources.Credits -= 5
			fmt.Println("ARCH PAID INFRA")
		}

	}

	for AllInfrastructure := 0; AllInfrastructure < len(d.Infrastructures); AllInfrastructure++ {
		if w.Government.Resources.Credits <= 0 {
			if d.Infrastructures[AllInfrastructure].ControlledBy == "" {
				d.Infrastructures[AllInfrastructure].State = "Maintenance"
				fmt.Println("ARCH LOST INFRA")
			}

		} else if d.Infrastructures[AllInfrastructure].ControlledBy == "" {
			w.Government.Resources.Credits -= 2
			fmt.Println("ARCH PAID INFRA")
		}

	}
}
