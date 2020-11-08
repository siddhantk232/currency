package main

import (
	"net"
	"os"

	"github.com/siddhantk232/currency/data"

	"google.golang.org/grpc/reflection"

	"github.com/hashicorp/go-hclog"
	"github.com/siddhantk232/currency/protos/currency"
	"github.com/siddhantk232/currency/server"

	"google.golang.org/grpc"
)

func main() {
	log := hclog.Default()
	gs := grpc.NewServer()

	rates, err := data.NewRates(log)

	if err != nil {
		log.Error("error getting rates", "error", err)
		os.Exit(1)
	}

	currencyHandler := server.NewCurrency(rates, log)

	currency.RegisterCurrencyServer(gs, currencyHandler)

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("unable to listen", "error", err)
		os.Exit(1)
	}

	gs.Serve(l)
}
