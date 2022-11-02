package order

import (
	"fmt"
	"github.com/rs/xid"
	"strings"
)

type Order struct {
	ID                       xid.ID
	Symbol                   string `json:"symbol"`
	ExchangeOrderId          int64  `json:"orderId"`
	Status                   int
	ClientOrderID            string  `json:"clientOrderId"`
	TransactTime             int64   `json:"transactTime"`
	Price                    string  `json:"price"`
	OrigQuantity             string  `json:"origQty"`
	ExecutedQuantity         string  `json:"executedQty"`
	CummulativeQuoteQuantity string  `json:"cummulativeQuoteQty"`
	StopLoss                 float64 `yaml:"stop_loss"`
	StopGain                 float64 `yaml:"stop_gain"`
	//RSIBuy      float64       `yaml:"rsi_buy"`
	//RSISell     float64       `yaml:"rsi_sell"`
	//RSILimit    int           `yaml:"rsi_limit"`
	RSIInterval    string `yaml:"rsi_interval"`
	ExchangeName   string
	Amount         float64
	PriceBeforeBuy float64
}

func NewOrder(symbol string, amount float64) (Order, error) {

	return Order{
		ID:          xid.New(),
		Status:      0,
		Symbol:      symbol,
		Amount:      amount,
		RSIInterval: "15m",
	}, nil
}

func (s *Order) GetAsset() string {
	return fmt.Sprintf("%s%s", s.GetSymbol(), s.BuyWith())
}

// GetSymbol split to get symbol that want to buy.
func (s *Order) GetSymbol() string {
	return strings.Split(s.Symbol, "/")[0]
}

// BuyWith split to get symbol used to buy.
func (s *Order) BuyWith() string {
	return strings.Split(s.Symbol, "/")[1]
}

//type Exchange struct {
//	ID xid.ID
//	Active bool
//	BaseSymbol string
//	apiKey string
//	apiSecret string
//	RSIInterval string
//}
