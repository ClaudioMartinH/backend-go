package types

type User struct {
	Id        string
	Firstname string
	Lastname  string
}

type Product struct {
	Id          string
	Name        string
	Description string
	Price       float64
}

type CartItem struct {
	Product  Product
	Quantity int
}

type CartData struct {
	Products   []Product
	TotalPrice float64
}
