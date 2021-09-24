//go:build wireinject
// +build wireinject

package wireexample

import "github.com/google/wire"
func InitializeEvent() Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}