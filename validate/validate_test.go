package validate_test

import (
	"testing"

	"fetch/api"
	"fetch/validate"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidateProcessReceiptRequest(t *testing.T) {
	testCases := []struct {
		name     string
		input    api.ProcessReceiptRequest
		expected bool
	}{
		{
			name: "valid request",
			input: api.ProcessReceiptRequest{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []api.Item{
					{
						ShortDescription: "Mountain Dew 12PK",
						Price:            "6.49",
					},
				},
				Total: "6.49",
			},
			expected: true,
		},
		{
			name: "invalid retailer",
			input: api.ProcessReceiptRequest{
				Retailer:     "T@rget",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []api.Item{
					{
						ShortDescription: "Mountain Dew 12PK",
						Price:            "6.49",
					},
				},
				Total: "6.49",
			},
			expected: false,
		},
		{
			name: "invalid purchase date",
			input: api.ProcessReceiptRequest{
				Retailer:     "Target",
				PurchaseDate: "20222-01-01",
				PurchaseTime: "13:01",
				Items: []api.Item{
					{
						ShortDescription: "Mountain Dew 12PK",
						Price:            "6.49",
					},
				},
				Total: "6.49",
			},
			expected: false,
		},
		{
			name: "invalid purchase time",
			input: api.ProcessReceiptRequest{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "113:01",
				Items: []api.Item{
					{
						ShortDescription: "Mountain Dew 12PK",
						Price:            "6.49",
					},
				},
				Total: "6.49",
			},
			expected: false,
		},
		{
			name: "not enough items",
			input: api.ProcessReceiptRequest{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Total:        "6.49",
			},
			expected: false,
		},
		{
			name: "invalid total",
			input: api.ProcessReceiptRequest{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []api.Item{
					{
						ShortDescription: "Mountain Dew 12PK",
						Price:            "6.49",
					},
				},
				Total: "6.499",
			},
			expected: false,
		},
		{
			name: "invalid item description",
			input: api.ProcessReceiptRequest{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []api.Item{
					{
						ShortDescription: "Mount@in Dew 12PK",
						Price:            "6.49",
					},
				},
				Total: "6.49",
			},
			expected: false,
		},
		{
			name: "invalid item price",
			input: api.ProcessReceiptRequest{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []api.Item{
					{
						ShortDescription: "Mountain Dew 12PK",
						Price:            "6.499",
					},
				},
				Total: "6.49",
			},
			expected: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := validate.ValidateProcessReceiptRequest(tc.input)
			if tc.expected {
				assert.Nil(t, got)
			} else {
				assert.NotNil(t, got)
			}
		})
	}
}

func TestValidateGetPointsRequest(t *testing.T) {
	testCases := []struct {
		name       string
		input      api.GetPointsRequest
		expected   bool
		expectedID uuid.UUID
	}{
		{
			name:       "valid uuid",
			input:      api.GetPointsRequest{ID: "7f617996-dcf1-478b-9ba2-ab45e0bb33e1"},
			expected:   true,
			expectedID: uuid.MustParse("7f617996-dcf1-478b-9ba2-ab45e0bb33e1"),
		},
		{
			name:       "invalid uuid",
			input:      api.GetPointsRequest{ID: "abc123-def321"},
			expected:   false,
			expectedID: uuid.Nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			id, got := validate.ValidateGetPointsRequest(tc.input)
			if tc.expected {
				assert.Nil(t, got)
			} else {
				assert.NotNil(t, got)
			}
			assert.Equal(t, tc.expectedID, id)
		})
	}
}
