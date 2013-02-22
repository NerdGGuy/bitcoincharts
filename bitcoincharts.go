package bitcoincharts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"log"
	"net/http"
	"time"
)

type Market struct {
	Symbol          string
	Currency        string
	Bid             float64
	Ask             float64
	Latest_Trade    int
	Open            float64
	High            float64
	Low             float64
	Close           float64
	Previous_Close  float64
	Volume          float64
	Currency_Volume float64
}

type BitcoinChartsError struct {
	MyError  string
	TheError string
}

func (e BitcoinChartsError) Error() string {
	return fmt.Sprintf("%s --- %s", e.MyError, e.TheError)
}

type BitcoinCharts struct {
	Timeout  time.Duration
	lasttime time.Time
	markets  []Market
}

func (bc *BitcoinCharts) getJSON() *BitcoinChartsError {
	if time.Since(bc.lasttime) < bc.Timeout {
		return nil
	}

	resp, err := http.Get("http://bitcoincharts.com/t/markets.json")

	if err != nil {
		//log.Fatal("err:", err)
		return &BitcoinChartsError{"Error requesting http://bitcoincharts.com/t/markets.json", err.Error()}
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		//log.Fatal("err:", err)
		return &BitcoinChartsError{"Error reading http://bitcoincharts.com/t/markets.json", err.Error()}
	}

	err = json.Unmarshal(body, &bc.markets)
	//fmt.Printf("%+v", bc.markets)
	if err != nil {
		//fmt.Println("error:", err)
	}

	bc.lasttime = time.Now()

	return nil
}

func (bc *BitcoinCharts) GetMarket(symbol string) (*Market, *BitcoinChartsError) {
	err := bc.getJSON()
	if err != nil {
		return nil, err
	}

	for _, market := range bc.markets {
		//fmt.Printf("test: %+v", market)
		switch market.Symbol {
		case symbol:
			return &market, nil
		}
	}
	return nil, nil
}

func (bc *BitcoinCharts) GetMarketsWithCurrency(currency string) (*[]Market, *BitcoinChartsError) {
	err := bc.getJSON()
	if err != nil {
		return nil, err
	}

	markets := []Market{}
	for _, market := range bc.markets {
		//fmt.Printf("test: %+v", market)
		switch market.Currency {
		case currency:
			markets = append(markets, market)
		}
	}
	return &markets, nil
}

func (bc *BitcoinCharts) GetMarketSymbols() (*[]string, *BitcoinChartsError) {
	err := bc.getJSON()
	if err != nil {
		return nil, err
	}

	list := make([]string, len(bc.markets))
	i := 0
	for _, market := range bc.markets {
		//fmt.Printf("test: %+v", market)
		list[i] = market.Symbol
		i++
	}
	return &list, nil
}
