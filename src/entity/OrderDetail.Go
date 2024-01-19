package entity

type OrderDetail struct {
    OrderID      int    `json:"order_id"`
    ConsumableID int    `json:"consumable_id"`
    Quantity     int    `json:"quantity"`
    Comment      string `json:"comment"`
}
