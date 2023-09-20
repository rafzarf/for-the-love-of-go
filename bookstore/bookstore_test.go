package bookstore_test

import (
	"bookstore"
	"testing"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:   "Spark Joy",
		Author:  "Marie Kondō",
		Copies:  2,
		Edition: "First Edition",
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:   "Spark Joy",
		Author:  "Marie Kondō",
		Copies:  2,
		Edition: "Second Edition",
	}
	want := 1
	result := bookstore.Buy(b)
	got := result.Copies
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestDiscount(t *testing.T) {
	// Create a test case with initial values.
	b := bookstore.Book{
		Title:             "Test Book",
		Author:            "Test Author",
		Copies:            5,
		Edition:           "First Edition",
		Discount:          10.0,
		RoyaltyPercentage: 0,
		IntialPrice:       20.0,
		FinalPrice:        0.0,
	}

	// Calculate the expected discounted price.
	expectedDiscount := b.IntialPrice * (b.Discount / 100)
	actualDiscount := b.IntialPrice - expectedDiscount
	// Call the Discount function to calculate the discount.
	discount, _ := bookstore.Discount(&b)

	// Check if the calculated discount matches the expected discount.
	if discount != actualDiscount {
		t.Errorf("Expected discount of %.2f, but got %.2f", actualDiscount, discount)
	}
}

func TestTotalPrice(t *testing.T) {
	// Create a test case with initial values.
	b := bookstore.Book{
		Title:             "Test Book",
		Author:            "Test Author",
		Copies:            5,
		Edition:           "First Edition",
		Discount:          10.0,
		RoyaltyPercentage: 0,
		IntialPrice:       20.0,
		FinalPrice:        0.0,
	}

	// Calculate the expected discounted price.
	expectedPrice := b.FinalPrice * float64(b.Copies)

	// Call the Discount function to calculate the discount.
	got := bookstore.TotalPrice(b)

	// Check if the calculated discount matches the expected discount.
	if got != expectedPrice {
		t.Errorf("Expected Price of %.2f, but got %.2f", expectedPrice, got)
	}
}
