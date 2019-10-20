package profit

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestMaxProfit1(t *testing.T) {
	stockPricesYesterday := []int{10, 7, 5, 8, 11, 9}
	calc := MakeProfitCalculator()
	log.Infof("Maximum profit that can be made is %d", calc.CalculateMaxProfit(stockPricesYesterday))
	require.Equal(t, calc.CalculateMaxProfit(stockPricesYesterday), 6)
}

func TestMaxProfit2(t *testing.T) {
	stockPricesYesterday := []int{10, 17, 7, 5, 8, 11, 9}
	calc := MakeProfitCalculator()
	log.Infof("Maximum profit that can be made is %d", calc.CalculateMaxProfit(stockPricesYesterday))
	require.Equal(t, calc.CalculateMaxProfit(stockPricesYesterday), 7)
}

func TestMaxProfit3(t *testing.T) {
	stockPricesYesterday := []int{10, 17, 7, 5, 8, 18, 9}
	calc := MakeProfitCalculator()
	log.Infof("Maximum profit that can be made is %d", calc.CalculateMaxProfit(stockPricesYesterday))
	require.Equal(t, calc.CalculateMaxProfit(stockPricesYesterday), 13)
}
