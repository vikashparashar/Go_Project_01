package main

import "fmt"

func main() {
	ns := GetNutritionalScore(NutritionalData{
		Energy:             EnergyFromKcal(0),
		Sugars:             SugarGram(10),
		SaturatedFattyAcid: SaturatedFattyAcid(2),
		Sodium:             SodiumMilligram(500),
		Fruits:             FruitsPercent(60),
		Fibre:              FiberGram(4),
		Protein:            ProteinGram(2),
	}, Food)

	fmt.Printf("Nutritional Score : %d", ns.Value)
	fmt.Printf("Nutri Score : %s", ns.GetNutriScore())

}
