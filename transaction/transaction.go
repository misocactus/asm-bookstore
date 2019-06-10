package transaction

import (
	"github.com/katesyberspace/asm-bookstore/book"
	"github.com/katesyberspace/asm-bookstore/user"
)

type Transaction struct {
	Buyer *user.User
	Seller *user.User
	Amount int
	Book	*book.Book
}