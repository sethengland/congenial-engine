package calculator_test

import (
	"testing"

	"fetch/calculator"
	"fetch/models"

	"github.com/stretchr/testify/assert"
)

func TestCalculator(t *testing.T) {
	testCases := []struct {
		name     string
		input    models.Receipt
		expected int
	}{
		{
			name: "example 1",
			input: models.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []models.Item{
					{
						ShortDescription: "Mountain Dew 12PK",
						Price:            "6.49",
					},
					{
						ShortDescription: "Emils Cheese Pizza",
						Price:            "12.25",
					},
					{
						ShortDescription: "Knorr Creamy Chicken",
						Price:            "1.26",
					},
					{
						ShortDescription: "Doritos Nacho Cheese",
						Price:            "3.35",
					},
					{
						ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
						Price:            "12.00",
					},
				},
				Total: "35.35",
			},

			expected: 28,
		},
		{
			name: "example 2",
			input: models.Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: "2022-03-20",
				PurchaseTime: "14:33",
				Items: []models.Item{
					{
						ShortDescription: "Gatorade",
						Price:            "2.25",
					},
					{
						ShortDescription: "Gatorade",
						Price:            "2.25",
					},
					{
						ShortDescription: "Gatorade",
						Price:            "2.25",
					},
					{
						ShortDescription: "Gatorade",
						Price:            "2.25",
					},
				},
				Total: "9.00",
			},

			expected: 109,
		},
		{
			name: "example 3: 1 item",
			input: models.Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: "2022-03-20",
				PurchaseTime: "14:33",
				Items: []models.Item{
					{
						ShortDescription: "Gatorade",
						Price:            "2.25",
					},
				},
				Total: "9.00",
			},

			expected: 99,
		},
		{
			name: "example 4 all conditions: 5 + 50 + 25 + 5 + 2 + 6 + 10 = 103",
			input: models.Receipt{
				Retailer:     "12345",
				PurchaseDate: "2022-03-11",
				PurchaseTime: "14:33",
				Items: []models.Item{
					{
						ShortDescription: "Gatorades",
						Price:            "10.00",
					},
					{
						ShortDescription: "Gatorade",
						Price:            "10.00",
					},
				},
				Total: "20.00",
			},

			expected: 103,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := calculator.CalculatePoints(tc.input)
			assert.Equal(t, tc.expected, got)
		})
	}
}
