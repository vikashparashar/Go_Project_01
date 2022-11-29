package main

import (
	"fmt"
	"time"

	"github.com/vikashparashar/Nutrition_Calculator/nut"
)

func main() {
	ns := nut.GetNutritionalScore(nut.NutritionalData{
		Energy:             nut.EnergyFromKcal(100),
		Sugars:             nut.SugarGram(10),         // considerd as negative values
		SaturatedFattyAcid: nut.SaturatedFattyAcid(2), // considerd as negative values
		Sodium:             nut.SodiumMilligram(500),  // considerd as negative values
		Fruits:             nut.FruitsPercent(60),     // considerd as positive values
		Fibre:              nut.FiberGram(4),          // considerd as positive values
		Protein:            nut.ProteinGram(2),        // considerd as positive values
	}, nut.Food)

	fmt.Println("_______________  Starting The Application  _______________")
	fmt.Print("\n\n We are calculating your score please wait")
	for i := 0; i <= 3; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Print(" .")
	}
	fmt.Println("\n\n ----------------------------------------------------- ")
	fmt.Println("                  Nutritional Calculator              ")
	fmt.Println("------------------------------------------------------")
	fmt.Printf("|               Nutritional Score : %d                 |\n", ns.Value)
	fmt.Printf("|                     Nutri Score : %s                 |\n", ns.GetNutriScore())
	fmt.Println(" ----------------------------------------------------- ")
	fmt.Println("                 *** End Of Programe ***              ")
	fmt.Println("")

}
