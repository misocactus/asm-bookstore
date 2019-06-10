package bookstore

import (
	"github.com/katesyberspace/asm-bookstore/transaction"
	"github.com/katesyberspace/asm-bookstore/user"
)

type BookstoreApp struct {
	Users []*user.User
	Transactions []*transaction.Transaction
}