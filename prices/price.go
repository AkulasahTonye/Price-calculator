package prices

import (
	"fmt"

	"github.com/Price-Calculator/Conversion"
	"github.com/Price-Calculator/fileManager"
)

type TaxPriceJob struct {
	TaskRate          float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func (job *TaxPriceJob) Process() {

	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taskPriceJob := price * (1 + job.TaskRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taskPriceJob)
	}
	job.TaxIncludedPrices = result
	fileManager.WriteJson(fmt.Sprintf("result_%0.f.json", job.TaskRate*100), job)
}

func NewTaxPriceJob(taxRate float64) *TaxPriceJob {
	return &TaxPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaskRate:    taxRate,
	}
}

func (job *TaxPriceJob) LoadData() {
	line, err := fileManager.ReadLines("price.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	prices, err := Conversion.StringsToFloats(line)

	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices

}
