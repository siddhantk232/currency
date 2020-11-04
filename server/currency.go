package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/siddhantk232/currency/protos/currency"
)

// Currency service
type Currency struct {
	log hclog.Logger
	currency.UnimplementedCurrencyServer
}

// NewCurrency creates new currency handler
func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{log: l}
}

// GetRate service registrar
func (c *Currency) GetRate(ctx context.Context, rr *currency.RateRequest) (*currency.RateResponse, error) {

	c.log.Info("handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())

	return &currency.RateResponse{Rate: 0.5}, nil

}
