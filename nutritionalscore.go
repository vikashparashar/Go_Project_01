package main

const (
	Food ScoreType = iota
	Beverage
	Water
	Cheese
)

type (
	EnergyKJ           float64
	SugarGram          float64
	SaturatedFattyAcid float64
	SodiumMilligram    float64
	FruitsPercent      float64
	FiberGram          float64
	ProteinGram        float64
)
type NutritionalData struct {
	Energy             EnergyKJ
	Sugars             SugarGram
	SaturatedFattyAcid SaturatedFattyAcid
	Sodium             SodiumMilligram
	Fruits             FruitsPercent
	Fibre              FiberGram
	Protein            ProteinGram
	IsWater            bool
}
type ScoreType int

type NutritionalScore struct {
	Value     int
	Positive  int
	Negative  int
	ScoreType ScoreType
}

var (
	energyLevels         = []float64{3350, 3015, 2680, 2345, 2010, 1675, 1340, 1005, 670, 335}
	sugerLevels          = []float64{45, 40, 36, 31, 27, 22.5, 18, 13.5, 9, 4.5}
	saturatedFattyAcid   = []float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	sodiumLevel          = []float64{900, 810, 720, 630, 540, 450, 360, 270, 180, 90}
	fiberLevels          = []float64{4.7, 3.7, 2.8, 1.9, 0.9}
	proteinLevels        = []float64{8, 6.4, 408, 3.2, 1.6}
	energyLevelsBeverage = []float64{270, 240, 210, 180, 150, 120, 90, 60, 30, 0}
	sugersLevelBeverage  = []float64{13.5, 12, 10.5, 9, 7.5, 6, 4.5, 3, 1.5, 0}
	scoreToLetter        = []string{"A", "B", "C", "D"}
)

func (e EnergyKJ) GetPoints(st ScoreType) int {
	if st == Beverage {

		return getPointsFromRange(float64(e), energyLevelsBeverage)
	}
	return getPointsFromRange(float64(e), energyLevels)
}

func (s SugarGram) GetPoints(st ScoreType) int {
	if st == Beverage {

		return getPointsFromRange(float64(s), sugersLevelBeverage)
	}
	return getPointsFromRange(float64(s), sugerLevels)
}

func (s SaturatedFattyAcid) GetPoints(st ScoreType) int {

	return getPointsFromRange(float64(s), saturatedFattyAcid)
}
func (s SodiumMilligram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(s), sodiumLevel)
}

func (f FruitsPercent) GetPoints(st ScoreType) int {
	if st == Beverage {
		if f > 80 {
			return 10
		} else if f > 60 {
			return 4
		} else if f > 40 {
			return 2
		}
		return 0
	}

	if f > 80 {
		return 5
	} else if f > 60 {
		return 2
	} else if f > 40 {
		return 1
	}
	return 0
}
func (f FiberGram) GetPoints(st ScoreType) int {

	return getPointsFromRange(float64(f), fiberLevels)
}
func (p ProteinGram) GetPoints(st ScoreType) int {

	return getPointsFromRange(float64(p), proteinLevels)
}

func GetNutritionalScore(n NutritionalData, st ScoreType) NutritionalScore {

	value := 0
	positive := 0
	negative := 0
	if st != Water {
		fruitPoints := n.Fruits.GetPoints(st)
		fibrePoints := n.Fibre.GetPoints(st)
		negative = n.Energy.GetPoints(st) + n.Sugars.GetPoints(st) + n.SaturatedFattyAcid.GetPoints(st) + n.Sodium.GetPoints(st)
		positive = fruitPoints + fibrePoints + n.Protein.GetPoints(st)

		if st == Cheese {
			value = negative - positive
		} else {
			if negative >= 11 && fruitPoints < 5 {
				value = negative - positive - fibrePoints
			} else {
				value = negative - positive
			}
		}
	}

	return NutritionalScore{
		Value:     value,
		Positive:  positive,
		Negative:  negative,
		ScoreType: st,
	}
}

func EnergyFromKcal(kcal float64) EnergyKJ {
	return EnergyKJ(kcal * 4.184)
}
func SodiumFromSalt(saltmg float64) SodiumMilligram {
	return SodiumMilligram(saltmg / 2.5)
}

func (ns NutritionalScore) GetNutriScore() string {
	if ns.ScoreType == Food {
		return scoreToLetter[getPointsFromRange(float64(ns.Value), []float64{18, 10, 2, -1})]
	}
	// scoreToLetter[1] = "A"
	// scoreToLetter[2] = "B"
	// scoreToLetter[3] = "C"
	// scoreToLetter[4] = "D"
	// scoreToLetter[5] = "E"
	if ns.ScoreType == Water {
		return scoreToLetter[0]

	}
	// if ns.ScoreType == Beverage{
	// return
	// }
	// if ns.ScoreType == Cheese{
	// 	return
	// }
	return scoreToLetter[getPointsFromRange(float64(ns.Value), []float64{9, 5, 1, -2})]
}
func getPointsFromRange(v float64, steps []float64) int {
	lenSteps := len(steps)
	for i, l := range steps {
		if v > l {
			return lenSteps - i
		}
	}
	return 0
}
