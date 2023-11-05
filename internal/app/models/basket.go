package models

type Basket struct {
	Id     int          `json:"id"`
	UserId int          `json:"user_id"`
	Items  []BasketItem `json:"basket_item"`
}
