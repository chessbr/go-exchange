package main

import (
	"time"

	"go.exchange/matchingengine/engine"
	"go.exchange/matchingengine/entities"
)

func main() {
	matchingEngine := engine.CreateEngine("IRBR3")

	engine.EnqueueOrder(matchingEngine, entities.Order{
		Date:      time.Now(),
		Id:        "order1",
		OrderType: entities.Buy,
		Quantity:  100,
		Price:     1540,
		OwnerId:   "user123",
	})

	engine.EnqueueOrder(matchingEngine, entities.Order{
		Date:      time.Now(),
		Id:        "order2",
		OrderType: entities.Sell,
		Quantity:  100,
		Price:     1540,
		OwnerId:   "user3",
	})
	engine.DebugOrders(matchingEngine)
}
