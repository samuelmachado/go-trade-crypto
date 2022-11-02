package order

import "github.com/pkg/errors"

var (
	// ErrEmptyExchange cannot be nil
	ErrEmptyExchange       = errors.New("repository cnanot be nil")
	ErrUnableToGetSymbol   = errors.New("unable to get symbol info")
	ErrUnableToCreateOrder = errors.New("unable to create order")
)
