package models

type BasketItem struct {
	Id        int     `json:"id"`
	ProductId int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	BasketId  int     `json:"basket_id"`
	Product   Product `json:"product"`
	Basket    Basket  `json:"basket"`
}
