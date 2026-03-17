package main

func decide(w *World, s *Sector, d *District, myFaction int) string {
	// Une faction peut choisir différente actions
	// - Attaquer
	// - Trade
	// - Ne rien faire
	// Pour ca il faut qu'elle est connaissance des elements à sa disposition
	// Pour connaitre les elements à sa disposition, il y a les crédits, data, influence
	// Scouting (connaitre ce que possède les autres)
	// Si elle veut une ressources elle peut soit TRADE pour l'avoir SOIT attaquer
	// Si les elements ne sont pas en sa faveur elle peut ne rien faire
	// En fonction de son choix il faut le retourner pour choisir l'action à venir
	data, credit, influence, scoutResultFaction := scouting(d, myFaction)
	delta := 0

	if data > 0 || credit > 0 || influence > 0 || d.Factions[scoutResultFaction].Alive == true {
		action := "trade"
		for mem := 0; mem < len(d.Factions[myFaction].Memory); mem++ {
			if d.Factions[myFaction].Memory[mem].Who == d.Factions[scoutResultFaction].Name {
				if d.Factions[scoutResultFaction].Members < d.Factions[myFaction].Members && d.Factions[myFaction].Resources.Credits > d.Factions[scoutResultFaction].Resources.Credits {
					action = "war"
					newFriend := w.WorldTimer - d.Factions[myFaction].Memory[mem].Age
					if newFriend > 50 {
						action = "trade"
						break
					}
					break
				}
			}
		}

		delta = 2 * d.Factions[myFaction].Resources.Credits

		if delta < d.Factions[scoutResultFaction].Resources.Credits {
			action = "war"
		}

		switch action {
		case "war":
			war(w, s, d, myFaction)
		case "trade":
			initTrade(w, d, myFaction)
		}

		return action
	}

	return "nothing"
}

func scouting(d *District, myFaction int) (int, int, int, int) {

	for scoutResultFaction := 0; scoutResultFaction < len(d.Factions); scoutResultFaction++ { // Pour chaque autre faction du district
		if myFaction != scoutResultFaction && d.Factions[scoutResultFaction].Alive == true { //Je m'exclue pour garder que les autres
			amountInfluence := 0
			amountCredits := 0
			amountData := 0
			if d.Factions[scoutResultFaction].Resources.Influence > d.Factions[myFaction].Resources.Influence {
				amountInfluence = d.Factions[scoutResultFaction].Resources.Influence - d.Factions[myFaction].Resources.Influence
			}
			if d.Factions[scoutResultFaction].Resources.Credits > d.Factions[myFaction].Resources.Credits {
				amountCredits = d.Factions[scoutResultFaction].Resources.Credits - d.Factions[myFaction].Resources.Credits
			}
			if d.Factions[scoutResultFaction].Resources.Data > d.Factions[myFaction].Resources.Data {
				amountData = d.Factions[scoutResultFaction].Resources.Data - d.Factions[myFaction].Resources.Data
			}
			return amountData, amountCredits, amountInfluence, scoutResultFaction

		}
	}

	return 0, 0, 0, 0
}
