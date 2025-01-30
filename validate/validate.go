package validate

import (
	"fmt"
	"regexp"
	"time"

	"fetch/api"

	"github.com/google/uuid"
)

func ValidateProcessReceiptRequest(req api.ProcessReceiptRequest) error {
	valid, err := regexp.MatchString("^[\\w\\s\\-&]+$", req.Retailer)
	if err != nil || !valid {
		return fmt.Errorf("retailer did not match regex")
	}

	_, err = time.Parse("2006-01-02", req.PurchaseDate)
	if err != nil {
		return fmt.Errorf("invalid purchaseDate format")
	}
	_, err = time.Parse("15:04", req.PurchaseTime)
	if err != nil {
		return fmt.Errorf("invalid purchaseTime format")
	}
	valid, err = regexp.MatchString("^\\d+\\.\\d{2}$", req.Total)
	if err != nil || !valid {
		return fmt.Errorf("receipt total did not match regex")
	}

	if len(req.Items) < 1 {
		return fmt.Errorf("there must be at least 1 item on the receipt")
	}
	for i, item := range req.Items {
		err = ValidateItem(item)
		if err != nil || !valid {
			return fmt.Errorf("items[%d] validation failed: %w", i, err)
		}
	}
	return nil
}

func ValidateItem(item api.Item) error {
	valid, err := regexp.MatchString("^[\\w\\s\\-]+$", item.ShortDescription)
	if err != nil || !valid {
		return fmt.Errorf("short description did not match regex")
	}

	valid, err = regexp.MatchString("^\\d+\\.\\d{2}$", item.Price)
	if err != nil || !valid {
		return fmt.Errorf("price did not match regex")
	}
	return nil
}

func ValidateGetPointsRequest(req api.GetPointsRequest) (uuid.UUID, error) {
	id, err := uuid.Parse(req.ID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid uuid format: %w", err)
	}
	return id, nil
}
