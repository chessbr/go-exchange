package engine

import "go.exchange/matchingengine/entities"

func findOrderInBook(orderBook []entities.Order, orderId string) (entities.Order, int) {
	// find in buyBook
	for i := 0; i < len(orderBook); i++ {
		if orderBook[i].Id == orderId {
			return orderBook[i], i
		}
	}
	return entities.Order{}, -1
}

func removeFromBook(orderBook []entities.Order, orderId string) ([]entities.Order, bool) {
	var _, index = findOrderInBook(orderBook, orderId)
	if index >= 0 {
		return append(orderBook[:index], orderBook[index+1:]...), true
	}
	return orderBook, false
}
