package calculator

import (
	"math"
	"strconv"
	"strings"
	"time"

	"fetch/models"
)

func CalculatePoints(receipt models.Receipt) int {
	points := 0

	for _, c := range receipt.Retailer {
		if isAlphaNumeric(byte(c)) {
			points += 1
		}
	}

	total := strings.Split(receipt.Total, ".")
	if total[1] == "00" {
		points += 50
	}

	if v, err := strconv.Atoi(total[1]); err == nil && v%25 == 0 {
		points += 25
	}

	points += (len(receipt.Items) / 2) * 5
	for _, item := range receipt.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 {
			f, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				continue
			}
			points += int(math.Ceil(f * 0.2))
		}
	}

	parsedDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if parsedDate.Day()%2 != 0 {
		points += 6
	}

	parsedTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	start := time.Date(0, 1, 1, 14, 0, 0, 0, time.UTC)
	end := time.Date(0, 1, 1, 16, 0, 0, 0, time.UTC)
	if parsedTime.After(start) && parsedTime.Before(end) {
		points += 10
	}

	return points
}

func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
