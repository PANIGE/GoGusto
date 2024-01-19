package entity

type Chart struct {
    ConsumableID int     `json:"consumable_id"`
    TruckID      int     `json:"truck_id"`
    Label        string  `json:"label"`
    Description  string  `json:"description"`
    Price        float64 `json:"price"`
}
