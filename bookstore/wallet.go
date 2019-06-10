package bookstore

type Wallet struct {
	Owner *User
	Transactions []*Transaction
}