package bookstore

type Transaction struct {
	Buyer *User
	Seller *User
	Amount int
	Book	*Book
}