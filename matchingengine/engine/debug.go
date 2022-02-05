package engine

import (
	"fmt"

	"go.exchange/matchingengine/entities"
)

func printBook(book []entities.Order) {
	fmt.Println("Price\tQty\tId")
	for i := 0; i < len(book); i++ {
		var order = book[i]
		fmt.Printf("%06d\t%03d\t%s\n", order.Price, order.Quantity, order.Id)
	}
}

func DebugOrders(engine *MatchingEngine) {
	fmt.Printf("Engine asset: %s\n", engine.asset)
	fmt.Println("")

	fmt.Println("=== BUY ===")
	printBook(engine.buyBook)

	fmt.Println("")

	fmt.Println("=== SELL ===")
	printBook(engine.sellBook)

	fmt.Println("")
}
