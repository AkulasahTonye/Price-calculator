package prices

import "fmt"

type TaxPriceJob struct {
	TaskRate          float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job TaxPriceJob) Process() {
	result := make(map[string]float64)
	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaskRate)
	}
	fmt.Println(result)
}

func NewTaxPriceJob(taxRate float64) *TaxPriceJob {
	return &TaxPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaskRate:    taxRate,
	}
}
