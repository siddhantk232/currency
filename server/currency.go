package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/siddhantk232/currency/data"
	"github.com/siddhantk232/currency/protos/currency"
)

// Currency service
type Currency struct {
	rates *data.ExchangeRates
	log   hclog.Logger
	currency.UnimplementedCurrencyServer
}

// NewCurrency creates new currency handler
func NewCurrency(rates *data.ExchangeRates, l hclog.Logger) *Currency {
	return &Currency{rates: rates, log: l}
}

// GetRate service registrar
func (c *Currency) GetRate(ctx context.Context, rr *currency.RateRequest) (*currency.RateResponse, error) {

	c.log.Info("handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())

	rate, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())

	if err != nil {
		return nil, err
	}

	return &currency.RateResponse{Rate: rate}, nil

}
