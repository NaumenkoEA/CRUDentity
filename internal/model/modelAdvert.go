package model

type Advert struct {
	ID      int     `json:"id" form:"id"`
	Address string  `json:"address" form:"address"`
	Price   float32 `json:"price" form:"price"`
}
