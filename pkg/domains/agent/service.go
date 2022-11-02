package agent

import (
	"context"
	"fmt"
	"github.com/samuelmachado/go-core/log"
	"github.com/samuelmachado/go-trade-crypto/pkg/domains/order"
	"strings"
)

type ServiceI interface {
	Buy(ctx context.Context, agent *Agent, symbol string, price float64) error
	Sell(ctx context.Context, agent *Agent, symbol string, price float64) error
}

type Service struct {
	OrderService order.ServiceI
	Log          log.Logger
}

func NewService(orderService order.ServiceI, log log.Logger) (*Service, error) {
	if orderService == nil {
		return nil, EmptyOrderService
	}

	return &Service{
		OrderService: orderService,
		Log:          log,
	}, nil
}

func (s *Service) Buy(ctx context.Context, ag *Agent, symbol string, price float64) error {

	s.Log.Info(ctx, "starting buying", log.Any("agent_id", ag.ID), log.Any("symbol", symbol))
	if ag.ActiveOrders >= ag.MaxActiveOrders {
		return ErrMaxActiveOrders
	}

	if !strings.HasSuffix(symbol, ag.BaseSymbol) {
		return ErrInvalidSymbol
	}

	amountQuota := (ag.Wallet / 100) * ag.QuotaPerOrder
	total := amountQuota / price

	s.Log.Info(ctx, "total selected for buy", log.Any("agent_id", ag.ID), log.Any("symbol", symbol), log.Any("total", total), log.Any("unit_price", price))

	od, err := order.NewOrder(symbol, total)
	if err != nil {
		return ErrUnableToCreateOrder
	}

	s.Log.Info(ctx, "order created", log.Any("agent_id", ag.ID), log.Any("symbol", symbol), log.Any("order_id", od.ID))

	err = s.OrderService.Buy(ctx, &od)
	if err != nil {
		return ErrUnableToExecuteOrder
	}
	s.Log.Info(ctx, "placed to execute", log.Any("agent_id", ag.ID), log.Any("symbol", symbol), log.Any("order_id", od.ID))

	ag.Orders = append(ag.Orders, od)
	s.Log.Info(ctx, "active orders", log.Any("agent_id", ag.ID), log.Any("orders", len(ag.Orders)))

	return nil
}

func (s *Service) Sell(ctx context.Context, ag *Agent, symbol string, price float64) error {
	fmt.Println(ag.Orders)
	return nil
}
