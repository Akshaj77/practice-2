package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/filemanager"
)

type TaxIncludedPricesJob struct {
	IOManager         filemanager.FileManager `json:"-"`
	TaxRate           float64                 `json:"tax_rate"`
	InputPrices       []float64               `json:"input_prices"`
	TaxIncludedPrices map[string]string       `json:"tax_included_prices"`
}

func (job *TaxIncludedPricesJob) LoadData() {

	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println(err)

		return
	}

	job.InputPrices = prices

}

func (job *TaxIncludedPricesJob) Process() {

	job.LoadData()
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	job.IOManager.WriteJSON(job)

}

func NewTaxIncludedPriceJob(taxRate float64, fileManger filemanager.FileManager) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
		IOManager:   fileManger,
	}
}
