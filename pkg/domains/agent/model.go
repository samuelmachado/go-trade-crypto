package agent

import (
	"github.com/dustinkirkland/golang-petname"
	"github.com/rs/xid"
	"github.com/samuelmachado/go-trade-crypto/pkg/domains/order"
)

type Agent struct {
	ID              xid.ID
	Name            string
	ActiveOrders    uint
	MaxActiveOrders uint
	Wallet          float64
	BaseSymbol      string
	Orders          []order.Order
	QuotaPerOrder   float64
}

const (
	MinSymbolLength  = 3
	MaxSymbolLength  = 4
	MaxQuotaPerOrder = 50
	MinQuotaPerOrder = 10
)

func NewAgent(maxActiveOrders uint, wallet float64, quotaPerOrder float64, baseSymbol string) (Agent, error) {
	if wallet <= 0.0 || maxActiveOrders < 1 || quotaPerOrder > MaxQuotaPerOrder || quotaPerOrder < MinQuotaPerOrder || len(baseSymbol) < MinSymbolLength || len(baseSymbol) > MaxSymbolLength {
		return Agent{}, ErrNew
	}
	return Agent{
		ID:              xid.New(),
		Name:            petname.Generate(2, "-"),
		BaseSymbol:      baseSymbol,
		QuotaPerOrder:   quotaPerOrder,
		Wallet:          wallet,
		MaxActiveOrders: maxActiveOrders,
	}, nil
}
