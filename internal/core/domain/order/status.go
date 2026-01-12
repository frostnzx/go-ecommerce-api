package order

type OrderStatus string

const (
	OrderPending	OrderStatus = "pending"
	OrderPaid		OrderStatus = "paid"
	OrderShipped	OrderStatus = "shipped"
	OrderCancelled	OrderStatus = "cancelled"
)
