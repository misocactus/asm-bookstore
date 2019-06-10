package bookstore

type User struct {
	Name string
	BooksOwned []*Book
	Wallet *Wallet
}

func CreateUser (name string) (user User) {
	return User{
		Name: name,
		BooksOwned: []*Book{},

	}
}