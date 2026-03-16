package main

func gathering(w *World, sector int, district int, faction int) {

	switch w.Sectors[sector].Districts[district].Factions[faction].Type {
	case "enterprise":
		w.Sectors[sector].Districts[district].Factions[faction].Resources.Credits += +w.Sectors[sector].Districts[district].Factions[faction].Strength
	case "collectif":
		w.Sectors[sector].Districts[district].Factions[faction].Resources.Data += +w.Sectors[sector].Districts[district].Factions[faction].Strength + w.Sectors[sector].Districts[district].Factions[faction].Resources.Credits
	case "mafia":
		w.Sectors[sector].Districts[district].Factions[faction].Resources.Influence += +w.Sectors[sector].Districts[district].Factions[faction].Strength * 3
	case "classified":
		w.Sectors[sector].Districts[district].Factions[faction].Resources.Credits += +w.Sectors[sector].Districts[district].Factions[faction].Strength * 2
	}
}
