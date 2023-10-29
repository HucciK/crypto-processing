package errors

import "errors"

var (
	ErrCantCreateObserver = errors.New("can't create observer for wallet")
	ErrEmptyCallResult    = errors.New("empty call result")
	ErrNoTransactionFound = errors.New("no transactions found")
)
