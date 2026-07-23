package checkout

import (
	"errors"
	"fmt"
)

type Item struct{
	Name string 
	Price float64 
}

type Order struct {
	ID string 
	Items []Item 
	PromoCode string 
	TaxRate float64 // 0.10 for 10%
}

var (
	ErrEmptyOrder  = errors.New("Order must contain at least one item!!")
	ErrInvalidTax  = errors.New("tax rate cannot be negative")
	ErrInvalidItem = errors.New("Item price must be greater than zero")
)

func CalculateTotal(order Order) (float64, error) {
	if len(order.Items) == 0 {
		return 0, ErrEmptyOrder
	}
	if order.TaxRate < 0 {
		return 0, ErrInvalidTax
	}
	var subtotal float64
	for _, item := range order.Items {
		if item.Price <= 0 {
			return 0, fmt.Errorf("%w: %s has non-positive price", ErrInvalidItem, item.Name)
		}
		subtotal += item.Price
	}
	// Apply discount codes 
	discount := 0.0
	switch order.PromoCode{
	case "MORNING10":
		discount = 10
	case "HALFPRICE":
		discount = subtotal*0.5 
	}

	subtotal -= discount 
	if subtotal < 0{
		subtotal = 0
	}
	total := subtotal* (1 + order.TaxRate)
	return total , nil
}