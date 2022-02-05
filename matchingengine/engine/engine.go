package engine

import (
	"go.exchange/matchingengine/entities"
)

type MatchingEngine struct {
	asset    string
	buyBook  []entities.Order
	sellBook []entities.Order
}

type IMatchingEngine interface {
	EnqueueOrder(engine *MatchingEngine, order entities.Order) entities.MatchingResult
	UpdateOrder(engine *MatchingEngine, order entities.Order) entities.MatchingResult
	DeleteOrder(engine *MatchingEngine, orderId string) bool
}

var engine *MatchingEngine = nil
