package main

// NikeFactory is a concrete factory that implements the ISportsFactory interface
type NikeFactory struct{}

func (n *NikeFactory) createShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "nike",
			size: 10,
		},
	}
}

func (n *NikeFactory) createShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "nike",
			size: 40,
		},
	}
}
