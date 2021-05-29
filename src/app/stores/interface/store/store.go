package store

import (
	"github.com/ngs/go-amazon-product-advertising-api/amazon"
)

//Operation type, every operation need Client param and return Result type
type Operation func(Client, ...map[string]string) Result

//Images type
type Images struct {
	Small  string
	Medium string
	Big    string
}

type Product struct {
	Id    string
	Title string
	Price string
	Store string
	Image []Images
}

//type ProductList struct {
//	Products []Product
//}

type Client struct {
	Client *amazon.Client
	Store  string
}

type Result struct {
	Result []Product
	Extra  string
	Total  int
	Pages  int
}

type IStore interface {
	GetStore() Client
	GetOperation() Operation
	GetParameters() map[string]string
	SearchItem(Client, ...map[string]string) Result
}
