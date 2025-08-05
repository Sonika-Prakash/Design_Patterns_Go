package main

// AdidasFactory is the concrete factory that implements the ISportsFactory interface
type AdidasFactory struct{}

func (a *AdidasFactory) createShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 11,
		},
	}
}

func (a *AdidasFactory) createShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "adidas",
			size: 41,
		},
	}
}
