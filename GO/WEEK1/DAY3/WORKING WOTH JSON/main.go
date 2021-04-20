package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Product struct {
	ProductId      int    `json:"productId"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"perPerUnit"`
	QuantityOnHand int    `json:"quantityOnHand"`
	ProductName    string `json:"productName"`
}

func main() {
	fmt.Println("hello ")
	product := &Product{
		ProductId:      123,
		Manufacturer:   "Big Box Company",
		Sku:            "12.99",
		Upc:            "4561qHJK",
		PricePerUnit:   "41423254356",
		QuantityOnHand: 28,
		ProductName:    "Gizmo",
	}
	product:= Product{}

	// productJSON, err := json.Marshal(product)
	// if err != nil {
	// 	log.Fatal(err)

	// }
	err := json.Unmarshal([]byte(productJSON), &product)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(productJSON))
	fmt.Println("json unmarshal product :%s", product.ProductName)

}
