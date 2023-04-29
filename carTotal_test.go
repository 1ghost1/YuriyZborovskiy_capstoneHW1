package main

import (
	"testing"
)

// No need to test incorrect type since function requests slice of specific type "Item", it is understood that such type of standard compile time errors dont need to be tested

func TestCalculateTotalEmptyList(t *testing.T) {

	_, err := CalculateTotal("de", []Item{})

	if err == nil { // We didnt encounter any errors so behavior is bad
		t.Error("ERROR! Function did not return error on empty array!")
	} else {
		t.Log("Caught error, this is expected, test passed.")
	}

}

func TestHappyCalculateTotal(t *testing.T) {

	happyCart := []Item{
		{"Test1", "WIC_Eligeble", 50.03},
		{"Cring Fur", "Clothing", 20.03},
		{"Test1 Fur", "Clothing", 20.03},
		{"Test1 Other", "Other", 4.03},
		{"Test1 Fffffueeeer", "Clothing", 20.03},
		{"Test1 ddddd Fur", "Other", 4.03},
		{"Test1 Other", "Other", 5.03},
		{"Chicken Nuggets", "WIC_Eligeble", 4.03},
		{"Chicken Bruhthers Frozen Chicken Wrap", "WIC_Eligeble", 15.03},
	}

	expectedTotals := make(map[string]float32)
	expectedTotals["de"] = 142.26999
	expectedTotals["nj"] = 145.77791
	expectedTotals["pa"] = 145.45898

	for state, expectedTotal := range expectedTotals {

		calculatedTotal, error := CalculateTotal(state, happyCart)
		if error != nil {
			t.Fatal("ERROR! Function returned error unexpectedy")
		}

		if expectedTotal != calculatedTotal {
			t.Fatalf("ERROR! Unexpected total; State %s; Calculated total: %f; Expected total: %f", state, calculatedTotal, expectedTotal)
		}

	}
	t.Log("All set")
}

func TestBadDataCalculateTotal(t *testing.T) {
	cartWithUnspecifiedType := []Item{
		{"Test1", "WIC_Eligeble", 50.03},
		{"Cring Fur", "Clothing", 20.03},
		{"Test1 Fur", "Clothing", 20.03},
		{"Test1 Other", "", 4.03},
		{"Test1 Fffffueeeer", "Clothing", 20.03},
		{"Test1 ddddd Fur", "", 4.03},
		{"Test1 Other", "Otdher", 5.03},
		{"Chicken Nuggets", "WIC_Eligeble", 4.03},
		{"Chicken Bruhthers Frozen Chicken Wrap", "WIC_Eligeble", 15.03},
	}

	_, error := CalculateTotal("de", cartWithUnspecifiedType)

	if error == nil {
		t.Error("ERROR! Function didnt return an error on unspecified item type!")
	}

	cartWithUnspecifiedType = []Item{
		{"Test1", "WIC_Eligeble", 50.03},
		{"Cring Fur", "Clothing", 20.03},
		{"Test1 Fur", "Clothing", -20.03},
		{"Test1 Other", "", 4.03},
		{"Test1 Fffffueeeer", "Clothing", 20.03},
		{"Test1 ddddd Fur", "", 4.03},
		{"Test1 Other", "Otdher", 5.03},
		{"Chicken Nuggets", "WIC_Eligeble", 4.03},
		{"Chicken Bruhthers Frozen Chicken Wrap", "WIC_Eligeble", 15.03},
	}

	_, error = CalculateTotal("de", cartWithUnspecifiedType)

	if error == nil {
		t.Error("ERROR! Function didnt return an error on negative price!")
	}

}
