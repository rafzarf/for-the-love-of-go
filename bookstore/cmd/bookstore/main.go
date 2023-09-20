package main

import (
	"bookstore"
	"fmt"
)

func main() {
	// FUNGSI BELI
	b := bookstore.Buy(bookstore.Book{
		Title:             "Minimalism",
		Author:            "Marie Kondo",
		Copies:            1,
		Edition:           "First Edition",
		Discount:          50.0,
		RoyaltyPercentage: 0,
		IntialPrice:       20000.0,
	})
	fmt.Printf("Title: %s\nAuthor: %s\nCopies: %d\nEdition: %s\nDiscount: %.2f%%\nRoyalty Percentage: %.2f%%\n",
		b.Title, b.Author, b.Copies, b.Edition, b.Discount, b.RoyaltyPercentage)

	// FUNGSI HARGA SETELAH DISKON
	priceAfterDiscount, intialValue := bookstore.Discount(&b)
	fmt.Printf("Final Price: %.2f, from the intial value as: %.2f\n", priceAfterDiscount, intialValue)

	// FUNGSI TOTAL HARGA YANG PERLU DIBAYAR
	totalPrice := bookstore.TotalPrice(b)
	fmt.Println("So the total Price is: ", totalPrice)
}
