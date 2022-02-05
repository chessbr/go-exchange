package engine

import "go.exchange/matchingengine/entities"

func UpdateOrder(engine *MatchingEngine, order entities.Order) entities.MatchingResult {
	// remove from queue and add to queue
	if DeleteOrder(engine, order.Id) {
		return EnqueueOrder(engine, order)
	}

	return entities.MatchingResult{
		ResultType: entities.Deleted,
		Quantity:   order.Quantity,
		Price:      order.Price,
	}
}
