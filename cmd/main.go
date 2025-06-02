package main

import (
	"github.com/Price-Calculator/prices"
)

func main() {

	taskRate := []float64{0, 0.1, 0.07, 0.15}

	for _, taskRate := range taskRate {
		priceJob := prices.NewTaxPriceJob(taskRate)
		priceJob.Process()
	}

}
