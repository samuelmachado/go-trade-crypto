package binance

import (
	"context"
	"github.com/adshao/go-binance/v2"
	sdk "github.com/adshao/go-binance/v2"
	"github.com/samuelmachado/go-core/log"
)

const (
	BinanceConfigPrefix = "BINANCE_"
)

type Config struct {
	ApiKey string `env:"API_KEY,required"`
	Secret string `env:"SECRET_KEY,required"`
}

// Binance represents this package.
type Binance struct {
	Log     log.Logger
	Context context.Context
	Client  *binance.Client
}

// NewBinance create a new binance instance.
func NewBinance(log log.Logger, cfg Config, ctx context.Context) (*Binance, error) {
	sdk.UseTestnet = true
	client := sdk.NewClient(cfg.ApiKey, cfg.Secret)

	return &Binance{
		Log:     log,
		Context: ctx,
		Client:  client,
	}, nil
}

func (me *Binance) GetName() string {
	return BinanceConfigPrefix[:len(BinanceConfigPrefix)-1]
}
