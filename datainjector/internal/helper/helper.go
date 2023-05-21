package helper

import (
	"math/rand"
	"time"
)

func RandomTime() time.Time {
	endDate := time.Now().UTC()
	startDate := time.Date(
		endDate.Year(), endDate.Month(), endDate.Day(),
		endDate.Hour()-1, 0, 0, 0,
		time.UTC,
	)

	duration := endDate.Sub(startDate)
	randomDuration := time.Duration(rand.Int63n(int64(duration)))
	return startDate.Add(randomDuration)
}

func Vegetables(index int) string {
	veggies := []string{
		"Cabbage",
		"Cauliflower",
		"Cucumber",
		"Corn",
		"Garlic",
		"Green pepper",
		"Hothouse cucumber",
		"Ice plant",
		"Jalapeno",
		"Kidney beans",
		"Lemon",
		"Lime",
		"Marjoram",
		"Melons",
		"Mushroom",
		"Mustard plant",
		"Nettles",
		"Onions",
		"Olives",
		"Oyster plant",
		"Peanut",
		"Potatoes",
		"Pumpkins",
		"Radish",
		"Red kidney beans",
		"Rosemary",
		"Sea lettuce",
		"Snow peas",
		"Soybeans",
		"Spaghetti squash",
		"Spinach",
		"Sweet corn",
		"Tabasco pepper",
		"Tomato",
		"Wasabi",
		"White eggplant",
		"Winter melon",
		"Yellow squash",
		"Zucchinie",
	}
	return veggies[index]
}
