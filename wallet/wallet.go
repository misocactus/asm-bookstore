package wallet

import (
	"github.com/katesyberspace/asm-bookstore/transaction"
	"github.com/katesyberspace/asm-bookstore/user"
)

type Wallet struct {
	Owner *user.User
	Transactions []*transaction.Transaction
}