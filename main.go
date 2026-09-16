package main

type Soldat struct {
	nom     string
	vie     int
	attaque int
}

var equipe = [6]Soldat{
	{"Arthas", 1200, 250},
	{"Kael", 850, 320},
	{"Thrall", 1500, 180},
	{"Sylvanas", 700, 400},
	{"Garrosh", 1000, 280},
	{"Jaina", 500, 450},
}

func afficherEquipe() {
	println("==== EQUIPE ====")
	println()
	for _, soldat := range equipe {
		println("Nom:", soldat.nom)
		println("Attaque:", soldat.attaque)
		println("Vie:", soldat.vie)
		println()
	}
}

func trouverPlusDeVie(equipe [6]Soldat) Soldat {
	soldatPlusDeVie := equipe[0]
	for _, soldat := range equipe {
		if soldat.vie > soldatPlusDeVie.vie {
			soldatPlusDeVie = soldat
		}
	}
	return soldatPlusDeVie
}

func trouverPlusDAttaque(equipe [6]Soldat) Soldat {
	soldatPlusDAttaque := equipe[0]
	for _, soldat := range equipe {
		if soldat.attaque > soldatPlusDAttaque.attaque {
			soldatPlusDAttaque = soldat
		}
	}
	return soldatPlusDAttaque
}

func calculerVieMoyenne(equipe [6]Soldat) float64 {
	totalVie := 0
	for _, soldat := range equipe {
		totalVie += soldat.vie
	}
	return float64(totalVie) / float64(len(equipe))
}

func compterFaibles(equipe [6]Soldat) int {
	compteur := 0
	for _, soldat := range equipe {
		if soldat.vie < 800 {
			compteur++
		}
	}
	return compteur
}

func main() {
	afficherEquipe()
	soldatPlusDeVie := trouverPlusDeVie(equipe)
	soldatPlusDAttaque := trouverPlusDAttaque(equipe)
	vieMoyenne := calculerVieMoyenne(equipe)
	println("==== ANALYSE ====")
	println()
	println("Soldat avec le plus de vie:")
	println("Nom:", soldatPlusDeVie.nom)
	println("Vie:", soldatPlusDeVie.vie)
	println()
	println("Soldat avec la plus grande attaque:")
	println("Nom:", soldatPlusDAttaque.nom)
	println("Attaque:", soldatPlusDAttaque.attaque)
	println()
	println("Vie moyenne:")
	println("Vie:", vieMoyenne)
	println()
	println("Soldats avec moins de 800 PV:")
	println("Nombre:", compterFaibles(equipe))

}
