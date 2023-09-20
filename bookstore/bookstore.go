package bookstore

// Book represents information about a book.
type Book struct {
	Title             string
	Author            string
	Copies            int
	Edition           string
	inStock           bool
	Discount          float64
	RoyaltyPercentage float64
	IntialPrice       float64
	FinalPrice        float64
}

func Buy(b Book) Book {
	b.Copies--
	return b
}

func Discount(b *Book) (float64, float64) {
	discount := b.IntialPrice * (b.Discount / 100)
	b.FinalPrice = b.IntialPrice - discount
	return b.FinalPrice, b.IntialPrice
}

func TotalPrice(b Book) float64 {
	totalPrice := b.FinalPrice * float64(b.Copies)
	return totalPrice
}
