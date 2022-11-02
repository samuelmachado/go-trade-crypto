package agent

import "github.com/pkg/errors"

var (
	// ErrNew represents an error when creating a new Agent
	ErrNew                  = errors.New("unable to create new agent")
	EmptyOrderService       = errors.New("order service cannot be nil")
	ErrMaxActiveOrders      = errors.New("the agent cannot buy any new orders ")
	ErrInvalidSymbol        = errors.New("you are trying to trade an invalid symbol for this agent")
	ErrUnableToCreateOrder  = errors.New("unable to create order")
	ErrUnableToExecuteOrder = errors.New("unable to execute order")
)
