package main

import (
	"context"
	"fmt"

	"fetch/api"
	"fetch/calculator"
	"fetch/models"
	"fetch/storage"
	"fetch/validate"
)

func ProcessReceipt(ctx context.Context, req api.ProcessReceiptRequest) (api.ProcessReceiptResponse, error) {
	err := validate.ValidateProcessReceiptRequest(req)
	if err != nil {
		return api.ProcessReceiptResponse{}, fmt.Errorf("invalid request: %w", err)
	}
	items := make([]models.Item, len(req.Items))
	for i, item := range req.Items {
		items[i] = models.Item{
			ShortDescription: item.ShortDescription,
			Price:            item.Price,
		}
	}
	receipt := models.Receipt{
		Retailer:     req.Retailer,
		PurchaseDate: req.PurchaseDate,
		PurchaseTime: req.PurchaseTime,
		Items:        items,
		Total:        req.Total,
	}
	points := calculator.CalculatePoints(receipt)
	id, err := storage.SaveRecord(ctx, points)
	if err != nil {
		return api.ProcessReceiptResponse{}, fmt.Errorf("error saving record to storage: %w", err)
	}
	fmt.Printf("saved receipt with id: %s, points: %d", id.String(), points)
	return api.ProcessReceiptResponse{ID: id.String()}, nil
}

func GetPoints(ctx context.Context, req api.GetPointsRequest) (api.GetPointsResponse, error) {
	id, err := validate.ValidateGetPointsRequest(req)
	if err != nil {
		return api.GetPointsResponse{}, fmt.Errorf("invalid request: %w", err)
	}

	points, err := storage.GetRecord(ctx, id)
	if err != nil {
		return api.GetPointsResponse{}, fmt.Errorf("error getting record from storage: %w", err)
	}
	return api.GetPointsResponse{Points: points}, nil
}
