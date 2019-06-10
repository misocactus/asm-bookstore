package bookstore_test

import (
	"github.com/katesyberspace/asm-bookstore/bookstore"
	"testing"
)

func TestCreateUser(t *testing.T) {
	testuser := bookstore.CreateUser("Ben")
	if testuser.Name != "Ben" {
		t.Errorf("Expected name to be Ben; got %s",testuser.Name)

	}
}
