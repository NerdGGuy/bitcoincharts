package bitcoincharts

import (
	"testing"
	//"log"
	"time"
)

var bitcoin BitcoinCharts = BitcoinCharts{Timeout: time.Millisecond * 1000}

func NewBitcoinCharts() *BitcoinCharts {
	return &BitcoinCharts{Timeout: time.Millisecond * 1000}
}

func GetMtgox(t *testing.T) *Market {
	mtgoxUSD, err := bitcoin.GetMarket("mtgoxUSD")
	if err != nil {
		t.Fatalf("%s: %s", err.myError, err.theError)
	}
	if mtgoxUSD.Latest_Trade == 0 {
		t.Fatal("mtgoxUSD.Latest_Trade == nil")
	}
	return mtgoxUSD
}

func TestTimeout(t *testing.T) {
	//bitcoin := NewBitcoinCharts()

	timer1 := time.Now()
	_ = GetMtgox(t)
	dur1 := time.Since(timer1)

	time.Sleep(time.Millisecond * 100)
	timer2 := time.Now()
	_ = GetMtgox(t)
	dur2 := time.Since(timer2)

	time.Sleep(time.Millisecond * 1000)
	timer3 := time.Now()
	_ = GetMtgox(t)
	dur3 := time.Since(timer3)

	if dur2 > (dur3/2) || dur2 > (dur1/2) {
		t.Fatalf("request1: %s\nrequest2 (cached): %s\nrequest3: %s", dur1, dur2, dur3)
	}
}

func TestGetSymbols(t *testing.T) {
	symbols, err := bitcoin.GetMarketSymbols()
	if err != nil {
		t.Fatalf("%s: %s", err.myError, err.theError)
	}

	for _, symbol := range *symbols {
		switch symbol {
		case "mtgoxUSD":
			return
		}
	}
	t.Fatal("symbol was not found: mtgoxUSD")
}

func TestGetMarketsWithCurrency(t *testing.T) {
	markets, err := bitcoin.GetMarketsWithCurrency("USD")
	if err != nil {
		t.Fatalf("%s: %s", err.myError, err.theError)
	}

	for _, market := range *markets {
		switch market.Symbol {
		case "mtgoxUSD":
			return
		}
	}
	t.Fatal("symbol was not found: mtgoxUSD")
}
