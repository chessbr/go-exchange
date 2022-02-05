package engine

func DeleteOrder(engine *MatchingEngine, orderId string) bool {
	var newBuyBook, removedBuy = removeFromBook(engine.buyBook, orderId)

	if removedBuy {
		engine.buyBook = newBuyBook
		return true
	}

	var newSellBook, removedSell = removeFromBook(engine.sellBook, orderId)

	if removedSell {
		engine.sellBook = newSellBook
		return true
	}

	return false
}
