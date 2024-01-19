package entity

import "time"

type Truck struct {
    ID         int       `json:"id"`
    Owner      int 	     `json:"owner"`
    Name       string    `json:"name"`
    SlotBuffer uint8     `json:"slot_buffer"`
    Opening    time.Time `json:"opening"`
    Closing    time.Time `json:"closing"`
}
