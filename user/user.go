package user

import (
	"github.com/katesyberspace/asm-bookstore/book"
	"github.com/katesyberspace/asm-bookstore/wallet"
)

type User struct {
	Name string
	BooksOwned []*book.Book
	Wallet *wallet.Wallet
}