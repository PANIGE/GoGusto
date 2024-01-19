package entity

import "time"

type Order struct {
    OrderID  int       `json:"order_id"`
    UserID   int       `json:"user_id"`
    TruckID  int       `json:"truck_id"`
    Time     time.Time `json:"time"`
}
