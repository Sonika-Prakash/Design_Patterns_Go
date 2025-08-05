package main

// ISportsFactory is the Abstract Factory interface.
// This has creation methods for abstract products: Shoes and Shirts.
type ISportsFactory interface {
	createShoe() IShoe
	createShirt() IShirt
}

func createFactory(brand string) ISportsFactory {
	switch brand {
	case "adidas":
		return &AdidasFactory{}
	case "nike":
		return &NikeFactory{}
	default:
		return nil
	}
}
