package main

import "fmt"

/*
Preparation au négociation d'un trade

Elles peuvent echouer en fonction de plusieurs facteurs
- Guerre récentes
- Faction alive ?
- En guerre avec ?
*/
func initTrade(w *World, d *District, faction int) {
	//Avant de lancer un trade, je verifie si la faction init a eut des griefs avec l'autre
	//En vérifie en fonction du temps actuel et de l'age de la derniere agression

	for otherFaction := 0; otherFaction < len(d.Factions); otherFaction++ {
		canTrade := true
		if otherFaction != faction && d.Factions[otherFaction].Alive == true {

			fmt.Println(d.Factions[faction].Name, "TRADE INIT", d.Factions[otherFaction].Name)

			for mem := 0; mem < len(d.Factions[otherFaction].Memory); mem++ {
				if d.Factions[otherFaction].Memory[mem].Who == d.Factions[faction].Name {
					// DISABLE FOR TESTING TRADE
					// tooSoon := w.WorldTimer - d.Factions[otherFaction].Memory[mem].Age
					// if tooSoon > 50 {
					// 	canTrade = false
					// 	break
					// }
				}
			}

			if canTrade && faction != otherFaction {
				switch d.Factions[faction].Type {
				case "enterprise":
					exchangeTo, exchangeAt := exchange(d.Factions[faction].Resources.Credits, d.Factions[otherFaction].Resources.Data, 50, 1, 10)

					if exchangeTo > 0 {

						d.Factions[otherFaction].Resources.Credits = d.Factions[otherFaction].Resources.Credits + exchangeTo
						d.Factions[faction].Resources.Credits = d.Factions[faction].Resources.Credits - exchangeTo

						d.Factions[otherFaction].Resources.Data = d.Factions[otherFaction].Resources.Data - exchangeAt
						d.Factions[faction].Resources.Data = d.Factions[faction].Resources.Data + exchangeAt

						fmt.Println("EXCHANGE enterprise", exchangeTo, exchangeAt)

					}
				case "collectif":

					exchangeTo, exchangeAt := exchange(d.Factions[faction].Resources.Data, d.Factions[otherFaction].Resources.Credits, 50, 1, 10)

					if exchangeTo > 0 {

						d.Factions[otherFaction].Resources.Data = d.Factions[otherFaction].Resources.Data + exchangeTo
						d.Factions[faction].Resources.Data = d.Factions[faction].Resources.Data - exchangeTo

						d.Factions[otherFaction].Resources.Credits = d.Factions[otherFaction].Resources.Credits - exchangeAt
						d.Factions[faction].Resources.Credits = d.Factions[faction].Resources.Credits + exchangeAt

						fmt.Println("EXCHANGE collectif", exchangeTo, exchangeAt)

					}
				case "mafia":
					exchangeTo, exchangeAt := exchange(d.Factions[faction].Resources.Influence, d.Factions[otherFaction].Resources.Credits, 50, 1, 10)

					if exchangeTo > 0 {

						d.Factions[otherFaction].Resources.Influence = d.Factions[otherFaction].Resources.Influence + exchangeTo
						d.Factions[faction].Resources.Influence = d.Factions[faction].Resources.Influence - exchangeTo

						d.Factions[otherFaction].Resources.Credits = d.Factions[otherFaction].Resources.Credits - exchangeAt
						d.Factions[faction].Resources.Credits = d.Factions[faction].Resources.Credits + exchangeAt

						fmt.Println("EXCHANGE mafia", exchangeTo, exchangeAt)

					}
				case "classified":
					exchangeTo, exchangeAt := exchange(d.Factions[faction].Resources.Credits, d.Factions[otherFaction].Resources.Data, 50, 1, 10)

					if exchangeTo > 0 {

						d.Factions[otherFaction].Resources.Credits = d.Factions[otherFaction].Resources.Credits + exchangeTo
						d.Factions[faction].Resources.Credits = d.Factions[faction].Resources.Credits - exchangeTo

						d.Factions[otherFaction].Resources.Data = d.Factions[otherFaction].Resources.Data - exchangeAt
						d.Factions[faction].Resources.Data = d.Factions[faction].Resources.Data + exchangeAt

						fmt.Println("EXCHANGE classified", exchangeTo, exchangeAt)
					}
				}
			}
		}
	}
}

func exchange(resourceGiven int, resourceTaken int, threshold int, ratio int, percent int) (int, int) {
	if resourceGiven > threshold {
		exchangeTo := resourceGiven / percent
		exchangeAt := exchangeTo * ratio

		if resourceTaken <= exchangeAt {
			return 0, 0
		}

		return exchangeTo, exchangeAt
	}

	return 0, 0
}

// func exchange(s *Sector, faction int, otherFaction int, threshold int, ratio int, percent int) (int, int) {
// 	if d.Factions[faction].Resources.Credits > threshold {
// 		exchangeTo := d.Factions[faction].Resources.Credits / percent
// 		exchangeAt := exchangeTo * ratio

// 		return exchangeTo, exchangeAt

// 	}

// 	return d.Factions[faction].Resources.Credits, d.Factions[otherFaction].Resources.Credits
// }
