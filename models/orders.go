package models

type Orders struct {
	OrderID     int    `json:"id"`
	Customer    string `json:"customer"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Gender      string `json:"gender"`
	OrderDate   string `json:"orderDate"`
	IsDelivered bool   `json:"isDelivered"`
}

type OrderDetails struct {
	ProductID int    `json:"id"`
	OrderID   int    `json:"orderId"`
	Product   string `json:"product"`
	Quantity  int    `json:"quantity"`
	Price     int    `json:"price"`
}
