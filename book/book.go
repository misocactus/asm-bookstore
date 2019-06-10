package book

import "github.com/katesyberspace/asm-bookstore/user"

type Book struct {
	Title string
	Owner *user.User
	ForSale bool
	SalePrice int
}