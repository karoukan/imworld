package main

import "fmt"

/*
Preparation au négociation d'un trade

Elles peuvent echouer en fonction de plusieurs facteurs
- Guerre récentes
- Faction alive ?
- En guerre avec ?
*/
func initTrade(w *World, s *Sector, faction int) {
	//Avant de lancer un trade, je verifie si la faction init a eut des griefs avec l'autre
	//En vérifie en fonction du temps actuel et de l'age de la derniere agression

	for otherFaction := 0; otherFaction < len(s.Factions); otherFaction++ {
		canTrade := true
		if otherFaction != faction && s.Factions[otherFaction].Alive == true {

			fmt.Println(s.Factions[faction].Name, "TRADE INIT", s.Factions[otherFaction].Name)

			for mem := 0; mem < len(s.Factions[otherFaction].Memory); mem++ {
				if s.Factions[otherFaction].Memory[mem].Who == s.Factions[faction].Name {
					// DISABLE FOR TESTING TRADE
					// tooSoon := w.WorldTimer - s.Factions[otherFaction].Memory[mem].Age
					// if tooSoon > 50 {
					// 	canTrade = false
					// 	break
					// }
				}
			}

			if canTrade {
				switch s.Factions[faction].Type {
				case "enterprise":
					exchangeTo, exchangeAt := exchange(s.Factions[faction].Resources.Credits, s.Factions[otherFaction].Resources.Data, 50, 1, 10)

					if exchangeTo > 0 {

						s.Factions[otherFaction].Resources.Credits = s.Factions[otherFaction].Resources.Credits + exchangeTo
						s.Factions[faction].Resources.Credits = s.Factions[faction].Resources.Credits - exchangeTo

						s.Factions[otherFaction].Resources.Data = s.Factions[otherFaction].Resources.Data - exchangeAt
						s.Factions[faction].Resources.Data = s.Factions[faction].Resources.Data + exchangeAt

						fmt.Println("EXCHANGE enterprise", exchangeTo, exchangeAt)

					}
				case "collectif":

					exchangeTo, exchangeAt := exchange(s.Factions[faction].Resources.Data, s.Factions[otherFaction].Resources.Credits, 50, 1, 10)

					if exchangeTo > 0 {

						s.Factions[otherFaction].Resources.Data = s.Factions[otherFaction].Resources.Data + exchangeTo
						s.Factions[faction].Resources.Data = s.Factions[faction].Resources.Data - exchangeTo

						s.Factions[otherFaction].Resources.Credits = s.Factions[otherFaction].Resources.Credits - exchangeAt
						s.Factions[faction].Resources.Credits = s.Factions[faction].Resources.Credits + exchangeAt

						fmt.Println("EXCHANGE collectif", exchangeTo, exchangeAt)

					}
				case "mafia":
					exchangeTo, exchangeAt := exchange(s.Factions[faction].Resources.Influence, s.Factions[otherFaction].Resources.Credits, 50, 1, 10)

					if exchangeTo > 0 {

						s.Factions[otherFaction].Resources.Influence = s.Factions[otherFaction].Resources.Influence + exchangeTo
						s.Factions[faction].Resources.Influence = s.Factions[faction].Resources.Influence - exchangeTo

						s.Factions[otherFaction].Resources.Credits = s.Factions[otherFaction].Resources.Credits - exchangeAt
						s.Factions[faction].Resources.Credits = s.Factions[faction].Resources.Credits + exchangeAt

						fmt.Println("EXCHANGE mafia", exchangeTo, exchangeAt)

					}
				case "classified":
					exchangeTo, exchangeAt := exchange(s.Factions[faction].Resources.Credits, s.Factions[otherFaction].Resources.Data, 50, 1, 10)

					if exchangeTo > 0 {

						s.Factions[otherFaction].Resources.Credits = s.Factions[otherFaction].Resources.Credits + exchangeTo
						s.Factions[faction].Resources.Credits = s.Factions[faction].Resources.Credits - exchangeTo

						s.Factions[otherFaction].Resources.Data = s.Factions[otherFaction].Resources.Data - exchangeAt
						s.Factions[faction].Resources.Data = s.Factions[faction].Resources.Data + exchangeAt

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
// 	if s.Factions[faction].Resources.Credits > threshold {
// 		exchangeTo := s.Factions[faction].Resources.Credits / percent
// 		exchangeAt := exchangeTo * ratio

// 		return exchangeTo, exchangeAt

// 	}

// 	return s.Factions[faction].Resources.Credits, s.Factions[otherFaction].Resources.Credits
// }
