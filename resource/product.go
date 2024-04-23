package resource

import (
    anquanssl "github.com/anquanssl/golang-sdk"
)

type Product struct {
    client *anquanssl.Client
}

func NewProduct(client *anquanssl.Client) *Product {
    return &Product{
        client: client,
    }
}

func (p *Product) ProductList() (map[string]interface{}, error) {
    return p.client.Get("/product/list", make(map[string]string), make(map[string]interface{}))
}
