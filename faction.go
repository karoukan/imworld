package main

func gathering(w *World, s int, i int) {

	switch w.Sectors[s].Factions[i].Type {
	case "enterprise":
		w.Sectors[s].Factions[i].Resources.Credits += +w.Sectors[s].Factions[i].Strength
	case "collectif":
		w.Sectors[s].Factions[i].Resources.Data += +w.Sectors[s].Factions[i].Strength + w.Sectors[s].Factions[i].Resources.Credits
	case "mafia":
		w.Sectors[s].Factions[i].Resources.Influence += +w.Sectors[s].Factions[i].Strength * 3
	case "classified":
		w.Sectors[s].Factions[i].Resources.Credits += +w.Sectors[s].Factions[i].Strength * 2
	}
}

// func exchange(w *World, s int, i int, stockResources int) {
// 	// Je veux echanger des ressources entre 2 factions en fonction d'un seuil atteint
// 	// Si j'ai trop de crédit alors j'echange avec une faction qui a un surplus d'autre chose
// 	// J'ai une faction qui donne et une qui recoit
// 	for otherFaction := 0; len(w.Sectors[otherFaction]); otherFaction++ {

// 	}

// 	// if w.Sectors[s].Factions[i].Resources.Data > stockResources {
// 	// }
// 	// if w.Sectors[s].Factions[i].Resources.Influence > stockResources {
// 	// }
// }
