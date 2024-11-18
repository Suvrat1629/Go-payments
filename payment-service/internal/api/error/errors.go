// grpc/errors.go (or wherever you want to define your errors)
package grpc

import "errors"

var (
	ErrInvalidAmount = errors.New("invalid payment amount")
	ErrDatabase      = errors.New("database error")
	ErrInternal      = errors.New("internal server error")
)