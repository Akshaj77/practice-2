package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {

	taxRates := []float64{0, 0.7, 0.1, 0.15}

	// result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		taxjob := prices.NewTaxIncludedPriceJob(taxRate, filemanager.New(
			"prices.txt",
			fmt.Sprintf("result%.0f.json", taxRate*100),
		))
		taxjob.Process()
	}

	// fmt.Println(result)
}
