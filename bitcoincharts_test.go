package bitcoincharts

import (
	"testing"
	//"log"
	"time"
)

func NewBitcoinCharts() *BitcoinCharts {
	return &BitcoinCharts{Timeout: time.Millisecond * 1000}
}

func GetMtgox(t *testing.T, bitcoin *BitcoinCharts) *Market {
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
	bitcoin := NewBitcoinCharts()

	timer1 := time.Now()
	_ = GetMtgox(t, bitcoin)
	dur1 := time.Since(timer1)

	time.Sleep(time.Millisecond * 100)
	timer2 := time.Now()
	_ = GetMtgox(t, bitcoin)
	dur2 := time.Since(timer2)

	time.Sleep(time.Millisecond * 1000)
	timer3 := time.Now()
	_ = GetMtgox(t, bitcoin)
	dur3 := time.Since(timer3)

	if dur2 > (dur3/2) || dur2 > (dur1/2) {
		t.Fatalf("request1: %s\nrequest2 (cached): %s\nrequest3: %s", dur1, dur2, dur3)
	}
}
