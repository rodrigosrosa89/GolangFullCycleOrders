package entity

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	return &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}, nil
}

func (o *Order) IsValid() error {
	return nil
}
