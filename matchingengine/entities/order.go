package entities

import "time"

type OrderTypeEnum int

const (
	Buy  OrderTypeEnum = 1
	Sell               = 2
)

type Order struct {
	Date      time.Time
	Id        string
	OrderType OrderTypeEnum
	Quantity  int
	Price     int
	OwnerId   string
}
