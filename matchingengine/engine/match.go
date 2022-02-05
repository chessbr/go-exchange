package engine

import "go.exchange/matchingengine/entities"

func matchOrder(engine *MatchingEngine, order entities.Order) (bool, entities.ExecutionResult) {
	return false, entities.ExecutionResult{}
}
