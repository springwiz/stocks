package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/springwiz/stocks/profit"
)

func main() {
	stockPricesYesterday := []int{10, 7, 5, 8, 11, 9}
	calc := profit.MakeProfitCalculator()
	log.Infof("Maximum profit that can be made is %d", calc.CalculateMaxProfit(stockPricesYesterday))
}
