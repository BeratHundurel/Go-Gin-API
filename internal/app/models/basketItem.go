package models

type BasketItem struct {
	Id        int     `json:"id"`
	ProductId int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Product   Product `json:"product"`
	BasketId  int     `json:"basket_id"`
	Basket    Basket  `json:"basket"`
}
