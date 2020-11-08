package data

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hashicorp/go-hclog"
)

// ExchangeRates with base EUR
type ExchangeRates struct {
	log   hclog.Logger
	rates map[string]float64
}

// NewRates get new ExchangeRates
func NewRates(l hclog.Logger) (*ExchangeRates, error) {
	er := &ExchangeRates{log: l, rates: map[string]float64{}}
	err := er.getRates()
	return er, err
}

// GetRate returns the ratio of two provided rates
func (er *ExchangeRates) GetRate(base, dest string) (float64, error) {
	baseRate, ok := er.rates[base]

	if !ok {
		return 0, fmt.Errorf("rate not found for currency %s", base)
	}

	destRate, ok := er.rates[dest]

	if !ok {
		return 0, fmt.Errorf("rate not found for currency %s", dest)
	}

	return destRate / baseRate, nil
}

// GetRates get new rates data from the eu bank servers
func (er *ExchangeRates) getRates() error {
	resp, err := http.DefaultClient.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("unexpected response from the data source")
	}

	defer resp.Body.Close()

	md := &Cubes{}
	xml.NewDecoder(resp.Body).Decode(md)

	for _, c := range md.CubeDate {
		r, err := strconv.ParseFloat(c.Rate, 64)
		if err != nil {
			return err
		}

		er.rates[c.Currency] = r
	}

	er.rates["EUR"] = 1

	return nil
}

// Cubes xml data from eu servers
type Cubes struct {
	CubeDate []Cube `xml:"Cube>Cube>Cube"`
}

// Cube xml data from eu servers
type Cube struct {
	Currency string `xml:"currency,attr"`
	Rate     string `xml:"rate,attr"`
}
