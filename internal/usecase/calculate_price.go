package usecase

import "br.com.simple_order_service/internal/entity"

type OrderInput struct {
	ID    string
	Price float64
	Tax   float64
}

type OrderOutput struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

// SOLID - D - Dependency Inversion Principle
type CalculateFinalPrice struct {
	OrderDBInterface entity.OrderDBInterface
}

func NewCalculateFinalPrice(orderDBInterface entity.OrderDBInterface) *CalculateFinalPrice {
	return &CalculateFinalPrice{
		OrderDBInterface: orderDBInterface,
	}
}

func (c *CalculateFinalPrice) Execute(input OrderInput) (*OrderOutput, error) {
	order, err := entity.NewOrder(input.ID, input.Price, input.Tax)
	if err != nil {
		return nil, err
	}

	err = order.CalculateFinalPrice()
	if err != nil {
		return nil, err
	}

	err = c.OrderDBInterface.Save(order)
	if err != nil {
		return nil, err
	}

	return &OrderOutput{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
