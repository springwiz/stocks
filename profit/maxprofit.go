package profit

import (
	log "github.com/sirupsen/logrus"
)

type ProfitCalculator struct {
}

func MakeProfitCalculator() *ProfitCalculator {
	return &ProfitCalculator{}
}

func (p *ProfitCalculator) CalculateMaxProfit(stockPricesYesterday []int) int {
	maxProfit := stockPricesYesterday[1] - stockPricesYesterday[0]
	minPrice := stockPricesYesterday[0]
	for indx := 1; indx < len(stockPricesYesterday); indx++ {
		if stockPricesYesterday[indx]-minPrice > maxProfit {
			maxProfit = stockPricesYesterday[indx] - minPrice
			log.Debugf("Current max profit is %d", maxProfit)
		}
		if stockPricesYesterday[indx] < minPrice {
			minPrice = stockPricesYesterday[indx]
			log.Debugf("Current min price is %d", minPrice)
		}
	}
	return maxProfit
}
