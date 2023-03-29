package test

import (
	"github.com/go-web-kits/dbx"
	. "github.com/onsi/ginkgo"
)

type H = map[string]interface{}

type Example struct {
	ID   uint   `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

var _ = Describe("xx", func() {
	It("", func() {
		dbx.Where(&Example{}, dbx.EQ{"id": 1, "name": "abc"})
	})
})
