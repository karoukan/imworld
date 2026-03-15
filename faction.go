package main

func gathering(w *World, s int, i int) {

	switch w.Sectors[s].Factions[i].Type {
	case "enterprise":
		w.Sectors[s].Factions[i].Resources.Credits += +w.Sectors[s].Factions[i].Strength
	case "collectif":
		w.Sectors[s].Factions[i].Resources.Data += +w.Sectors[s].Factions[i].Strength + w.Sectors[s].Factions[i].Resources.Credits
	case "mafia":
		w.Sectors[s].Factions[i].Resources.Credits += +w.Sectors[s].Factions[i].Strength
	}
}
