package checkout

import (
	"testing"
)

func TestCalculateTotal_Basic(t *testing.T){
	order := Order{
		ID: "ORD-101",
		Items: []Item{{Name: "Latte" , Price : 5.00}},
		TaxRate: 0.10,
	}

	got,err := CalculateTotal(order)
	if err != nil {
		t.Fatalf("unexpected error : %v" , err)
	}

	want := 5.50 
	if got != want {
		t.Errorf("got %.2f , want %.2f" , got , want)
	}
}

func TestCalculateTotal_TableDriven(t *testing.T) {
	// Define our test cases table
	tests := []struct {
		name      string
		order     Order
		wantTotal float64
		wantErr   error
	}{
		{
			name: "Standard order with no promo",
			order: Order{
				Items:   []Item{{Name: "Espresso", Price: 3.00}, {Name: "Muffin", Price: 4.00}},
				TaxRate: 0.10,
			},
			wantTotal: 7.70, // (3 + 4) * 1.10
			wantErr:   nil,
		},
		{
			name: "Fixed discount promo (MORNING10)",
			order: Order{
				Items:     []Item{{Name: "Coffee Bag", Price: 20.00}},
				PromoCode: "MORNING10",
				TaxRate:   0.05,
			},
			wantTotal: 10.50, // (20 - 10) * 1.05
			wantErr:   nil,
		},
		{
			name: "Percentage promo (HALFPRICE)",
			order: Order{
				Items:     []Item{{Name: "Pour Over", Price: 10.00}},
				PromoCode: "HALFPRICE",
				TaxRate:   0.10,
			},
			wantTotal: 5.50, // (10 * 0.5) * 1.10
			wantErr:   nil,
		},
		{
			name: "Error on empty order",
			order: Order{
				Items:   []Item{},
				TaxRate: 0.10,
			},
			wantTotal: 0,
			wantErr:   ErrEmptyOrder,
		},
		{
			name: "Error on negative tax rate",
			order: Order{
				Items:   []Item{{Name: "Cold Brew", Price: 4.50}},
				TaxRate: -0.05,
			},
			wantTotal: 0,
			wantErr:   ErrInvalidTax,
		},
	}

	for _, tt := range tests {
		// t.Run isolates each row as its own subtest
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateTotal(tt.order)

			// Checking error handling
			if tt.wantErr != nil {
				if err == nil {
					t.Fatalf("expected error %v, got nil", tt.wantErr)
				}
				return // Stop here for error test cases
			}

			// If no error expected, ensure got is clean
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != tt.wantTotal {
				t.Errorf("got total %.2f; want %.2f", got, tt.wantTotal)
			}
		})
	}
}