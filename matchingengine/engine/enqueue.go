package engine

import "go.exchange/matchingengine/entities"

func EnqueueOrder(engine *MatchingEngine, order entities.Order) entities.MatchingResult {
	var matched, execution = matchOrder(engine, order)

	if matched {
		return entities.MatchingResult{
			ResultType: entities.Executed,
			Quantity:   execution.Quantity,
			Price:      execution.Price,
		}
	}

	if order.OrderType == entities.Buy {
		engine.buyBook = append(engine.buyBook, order)
	} else {
		engine.sellBook = append(engine.sellBook, order)
	}

	return entities.MatchingResult{
		ResultType: entities.Queued,
		Quantity:   order.Quantity,
		Price:      order.Price,
	}
}
