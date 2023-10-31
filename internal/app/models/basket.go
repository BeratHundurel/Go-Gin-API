package models

type Basket struct {
	Id     int
	UserId int
	Items  []BasketItem
}