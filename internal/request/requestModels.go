package request

type AdvertCreate struct {
	ID      int     `json:"id" form:"id"`
	Address string  `json:"address" form:"address"`
	Price   float32 `json:"price" form:"price"`
}

type AdvertUpdate struct {
	Address string  `json:"address" form:"address"`
	Price   float32 `json:"price" form:"price"`
}
