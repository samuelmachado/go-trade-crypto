package order

import (
	"context"
	"fmt"
	"github.com/samuelmachado/go-core/log"
)

type ExchangeI interface {
	CreateOrder(symbol, buyWith, side string, quantity, price float64) (err error)
	SymbolBalance(symbol string) (float64, error)
	SymbolPrice(symbol, interval string) (float64, error)
	GetName() string
}

type ServiceI interface {
	Buy(ctx context.Context, order *Order) error
	Sell(ctx context.Context, order *Order) error
	GetName() string
}

type Service struct {
	Exchange ExchangeI
	Log      log.Logger
}

func (s Service) GetName() string {
	return s.Exchange.GetName()
}

func NewService(exchange ExchangeI, log log.Logger) (*Service, error) {
	if exchange == nil {
		return nil, ErrEmptyExchange
	}

	return &Service{
		Exchange: exchange,
		Log:      log,
	}, nil
}

func (s *Service) Buy(ctx context.Context, order *Order) error {

	price, err := s.Exchange.SymbolPrice(order.GetAsset(), order.RSIInterval)

	if err != nil {
		s.Log.Error(ctx, ErrUnableToGetSymbol.Error(), log.Error(err), log.Any("order", order.ID))
		return ErrUnableToGetSymbol
	}

	order.PriceBeforeBuy = price

	fmt.Printf("%+v", order)
	err = s.Exchange.CreateOrder(order.GetSymbol(), order.BuyWith(), "BUY", order.Amount, price)
	if err != nil {
		s.Log.Error(
			ctx,
			ErrUnableToCreateOrder.Error(),
			log.Error(err),
			log.Any("order", order.ID),
			log.Any("operation", "BUY"),
		)
		return ErrUnableToCreateOrder
	}

	return nil
}

func (s *Service) Sell(ctx context.Context, order *Order) error {
	price, err := s.Exchange.SymbolPrice(order.GetAsset(), order.RSIInterval)

	if err != nil {
		s.Log.Error(ctx, ErrUnableToGetSymbol.Error(), log.Error(err), log.Any("order", order.ID))
		return ErrUnableToGetSymbol
	}

	order.PriceBeforeBuy = price

	err = s.Exchange.CreateOrder(order.GetSymbol(), order.BuyWith(), "SELL", order.Amount, price)
	if err != nil {
		s.Log.Error(
			ctx,
			ErrUnableToCreateOrder.Error(),
			log.Error(err),
			log.Any("order", order.ID),
			log.Any("operation", "SELL"),
		)
		return ErrUnableToCreateOrder
	}

	return nil
}
