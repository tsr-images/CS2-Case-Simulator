package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type Item struct {
	Rarity    string
	Condition string
	Float     float64

	//Not technically part of the Item struct, it should be tied to Item Quality
	Color string
}

var (
	inventory  []Item
	rarityName string
	condition  string

	// Colors
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"
)

var rarityColors = map[string]string{
	"Mil-Spec":           Cyan,    // Blue
	"Restricted":         Blue,    // Purple
	"Classified":         Magenta, // Pink
	"Covert":             Red,     // Red
	"★ Exceedingly Rare": Yellow,  // Gold
}

// func writeCsv(inv []Item) {
// 	var data [][]string

// 	file, err := os.OpenFile("inventory.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		log.Fatalln("Failed to write to file", err)
// 	}
// 	defer file.Close()

// 	w := csv.NewWriter(file)
// 	defer w.Flush()

// 	for _, entry := range inv {
// 		convertedFloat := fmt.Sprintf("%.12f", entry.Float)

// 		row := []string{convertedFloat, entry.Condition, entry.Rarity}
// 		data = append(data, row)
// 	}

// 	w.WriteAll(data)
// }

func main() {
	var n, err = strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)

	} else if n < 1 {
		fmt.Println("Error: Value cannot be less than 1")
	} else {
		for i := 0; i < n; i++ {

			floatValue := rand.Float64()
			rarity := rand.Float64()

			switch {
			case floatValue >= 0.45:
				condition = "Battle-Scarred"
			case floatValue >= 0.38:
				condition = "Well-Worn"
			case floatValue >= 0.15:
				condition = "Field-Tested"
			case floatValue >= 0.07:
				condition = "Minimal Wear"
			case floatValue >= 0:
				condition = "Factory New"
			}

			switch {
			case rarity <= 0.7992:
				rarityName = "Mil-Spec"
			case rarity <= 0.7992+0.1598:
				rarityName = "Restricted"
			case rarity <= 0.7992+0.1598+0.032:
				rarityName = "Classified"
			case rarity <= 0.7992+0.1598+0.032+0.0064:
				rarityName = "Covert"
			case rarity <= 0.7992+0.1598+0.032+0.0064+0.0026:
				rarityName = "★ Exceedingly Rare"
			}

			c := rarityColors[rarityName]
			newItem := Item{rarityName, condition, floatValue, c}
			inventory = append(inventory, newItem)
		}

		for _, item := range inventory {
			fmt.Printf(item.Color+"%.12f\t%s\t%s\t\n"+Reset, item.Float, item.Condition, item.Rarity)
		}

		//TODO
		// if os.Args[2] == "--csv" {
		// 	writeCsv(inventory)
		// }
	}
}
