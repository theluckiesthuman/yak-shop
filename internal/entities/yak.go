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
		if ageInDays >= 1000 {
			return totalMilk
		}

		dailyMilk := 50 - ageInDays*0.03
		if dailyMilk > 0 {
			totalMilk += dailyMilk
		}
	}
	return totalMilk
}

func (y *Yak) GetSkinStockTillDay(elapsedDays int) int {
	ageInDay := int(y.Age * 100)
	totalSkinInStock := 0

	for day := 0; day < elapsedDays; day++ {
		if y.shouldShaveToday(ageInDay+day, day) {
			totalSkinInStock++
		}
	}

	return totalSkinInStock
}

func (y *Yak) shouldShaveToday(currentAge, day int) bool {
	if !y.IsFirstTimeShaved {
		y.IsFirstTimeShaved = true
		return true
	} else {
		return day > (8 + int(float64(currentAge)*0.01))
	}
}

func (y *Yak) CalculateLastShavedAge(elapsedDays int) float64 {
	ageInDay := int(y.Age * 100)
	for day := 0; day < elapsedDays; day++ {
		if y.shouldShaveToday(ageInDay+day, day) {
			y.AgeLastShaved = y.Age + float64(day)/100
		}
	}
	return y.AgeLastShaved
}
