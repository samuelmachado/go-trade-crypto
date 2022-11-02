package container

import (
	"context"
	"github.com/samuelmachado/go-core/env"
	"github.com/samuelmachado/go-core/log"
	"github.com/samuelmachado/go-trade-crypto/pkg/domains/agent"
	"github.com/samuelmachado/go-trade-crypto/pkg/domains/order"

	"github.com/samuelmachado/go-trade-crypto/pkg/core/binance"
)

// Components are a like service, but it doesn't include business case
type components struct {
	Log     log.Logger
	Binance order.ExchangeI
}

// Services hold the business case, and make the bridge between
type Services struct {
	Order order.ServiceI
}

type Dependency struct {
	Components components
	Services   Services
}

func New(ctx context.Context) (context.Context, *Dependency, error) {
	cmp, err := setupComponents(ctx)
	if err != nil {
		return nil, nil, err
	}

	orderService, err := order.NewService(cmp.Binance, cmp.Log)
	//botService, err := bot.NewBot(cmp.Log, ctx, cmp.Binance, configs)
	if err != nil {
		return nil, nil, err
	}

	//One agent
	agentService, err := agent.NewService(orderService, cmp.Log)
	if err != nil {
		return nil, nil, err
	}

	ag, _ := agent.NewAgent(1, 1000, 25, "USDT")
	err = agentService.Buy(ctx, &ag, "BTC/USDT", 43810.8)
	if err != nil {
		cmp.Log.Error(ctx, "error on buying", log.Error(err), log.Any("agent_id", ag.ID))
	}
	srv := Services{
		Order: orderService,
	}

	dep := Dependency{
		Components: *cmp,
		Services:   srv,
	}

	return ctx, &dep, err

}

func setupComponents(ctx context.Context) (*components, error) {

	logInstance, err := log.NewLoggerZap(log.ZapConfig{
		Version:           "v0.1.0",
		DisableStackTrace: false,
	})
	if err != nil {
		return nil, err
	}

	binanceConfig := binance.Config{}

	err = env.LoadEnv(ctx, &binanceConfig, binance.BinanceConfigPrefix)
	if err != nil {
		return nil, err
	}

	binanceInstance, err := binance.NewBinance(logInstance, binanceConfig, ctx)
	if err != nil {
		return nil, err
	}

	return &components{
		Log:     logInstance,
		Binance: binanceInstance,
	}, nil
}
