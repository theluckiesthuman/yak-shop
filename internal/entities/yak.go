package entities

type YakTribe string

const (
	LabYaks YakTribe = "LabYaks"
)

type Herd struct {
	Yaks []Yak
}

type Yak struct {
	ID                int
	Name              string
	Age               float64
	Sex               string
	Tribe             YakTribe
	IsFirstTimeShaved bool
	AgeLastShaved     float64
}

func (y Yak) GetCurrentYakAgeInYears(elapsedDays int) float64 {
	return y.Age + float64(elapsedDays)/100
}

func (y Yak) CalculateMilk(elapsedDays int) float64 {
	totalMilk := 0.0
	for day := 0; day < elapsedDays; day++ {
		ageInDays := y.Age*100 + float64(day)

		// A LabYak dies the day it turns 10.
		if y.isDead(ageInDays) {
			return totalMilk
		}

		dailyMilk := 50 - ageInDays*0.03
		if dailyMilk > 0 {
			totalMilk += dailyMilk
		}
	}
	return totalMilk
}

func (y Yak) isDead(ageInDays float64) bool {
	return ageInDays >= 1000
}

func (y *Yak) GetSkinStockTillDay(elapsedDays int) int {
	totalSkinInStock := 0

	for day := 0; day < elapsedDays; day++ {
		ageInDays := y.Age*100 + float64(day)

		// A LabYak dies the day it turns 10.
		if y.isDead(ageInDays) {
			return totalSkinInStock
		}
		if y.shouldShaveToday(ageInDays, day) {
			totalSkinInStock++
		}
	}

	return totalSkinInStock
}

func (y *Yak) shouldShaveToday(currentAge float64, day int) bool {
	if !y.IsFirstTimeShaved {
		y.IsFirstTimeShaved = true
		return true
	} else {
		return float64(day) > float64(8+currentAge*0.01)
	}
}

func (y *Yak) CalculateLastShavedAge(elapsedDays int) float64 {
	for day := 0; day < elapsedDays; day++ {
		ageInDays := y.Age*100 + float64(day)

		if y.isDead(ageInDays) {
			break
		}
		if y.shouldShaveToday(ageInDays, day) {
			y.AgeLastShaved = y.Age + float64(day)/100
		}
	}
	return y.AgeLastShaved
}
