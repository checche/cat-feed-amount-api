package models

import "math"

type Cat struct {
	Weight     float64
	Age        int
	IsNeutered bool
	Feed       Feed
}

func (cat *Cat) calcFeedKcal() float64 {
	RER := 70 * math.Pow(cat.Weight, 0.75)
	var factor float64
	switch {
	case cat.Age < 4:
		factor = 3
	case cat.Age >= 4 && cat.Age < 7:
		factor = 2.5
	case cat.Age >= 7 && cat.Age < 12:
		factor = 2
	default:
		if cat.IsNeutered {
			factor = 1.2
		} else {
			factor = 1.4
		}
	}

	return RER * factor

}

func (cat *Cat) CalcFeedAmount() (float64, float64) {
	feedKcal := cat.calcFeedKcal()
	kcalPerUnit := cat.Feed.Kcal / float64(cat.Feed.PerGram)
	return feedKcal / kcalPerUnit, feedKcal
}
