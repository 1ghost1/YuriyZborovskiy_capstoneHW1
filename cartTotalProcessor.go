package main

import (
	"errors"
	"fmt"
	"strings"
)

type Item struct {
	Name  string
	Type  string
	Price float32
}

//var ALLOWED_ITEM_TYPES = []string{"WIC_Eligeble", "Clothing", "Other"} dont think i will actually need that, left for reference. Default in second switch basically handles unwanted items since if they are unwanted or unexpected we wont have implementation for that

func main() {
	//items := []Item{Item{"Test1", "WIC_Eligeble", 50.03}, Item{"Cring Fur", "Clothing", 20.03}, Item{"Test1 Fur", "Clothing", 20.03}, Item{"Test1 Other", "Other", 4.03}}
	happyCart := []Item{
		Item{"Test1", "WIC_Eligeble", 50.03},
		Item{"Cring Fur", "Clothing", 20.03},
		Item{"Test1 Fur", "Clothing", 20.03},
		Item{"Test1 Other", "Other", 4.03},
		Item{"Test1 Fffffueeeer", "Clothing", 20.03},
		Item{"Test1 ddddd Fur", "Other", 4.03},
		Item{"Test1 Other", "Other", 5.03},
		Item{"Chicken Nuggets", "WIC_Eligeble", 4.03},
		Item{"Chicken Bruhthers Frozen Chicken Wrap", "WIC_Eligeble", 15.03},
	}
	total, _ := CalculateTotal("pa", happyCart)

	fmt.Println(total)
}

func CalculateTotal(stateCode string, itemCart []Item) (float32, error) {

	total := float32(0.0)
	furTaxable := false
	var taxRate float32
	stateCode = strings.ToLower(stateCode)
	switch stateCode {
	case "de":
		taxRate = 0.0
		break
	case "nj":
		taxRate = 0.066
		furTaxable = true
		break
	case "pa":
		furTaxable = true
		taxRate = 0.06
		break

	default:
		return -1.0, errors.New("unsupported_state")

	}

	if len(itemCart) != 0 {
		fmt.Println(taxRate)
		for _, item := range itemCart {

			if item.Price < 0.0 {
				return -1, errors.New("negative_price")
			}

			fmt.Println(item)
			valueOfTax := float32(0.0)
			switch item.Type {
			case "WIC_Eligeble":
				valueOfTax = 0.0

			case "Clothing":
				if furTaxable && strings.Contains(strings.ToLower(item.Name), "fur") {
					valueOfTax = float32(item.Price) * taxRate
				} else {
					valueOfTax = 0.0
				}

			case "Other":
				valueOfTax = float32(item.Price) * taxRate

			default:
				return -1, errors.New("unsupported_item_type")
			}

			total += float32(item.Price) + float32(valueOfTax)
			fmt.Println("TAX ", taxRate, valueOfTax, " PRICE ", item.Price, " ITEM PRICE ", float32(item.Price)+float32(valueOfTax), " RUNNING TOTAL ", total)

		}

	} else {
		return -1.0, errors.New("passed_empty_cart") // no empty carts allowed
	}

	return float32(total), nil
}
