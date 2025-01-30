package main

import (
	"fmt"
	"net/http"

	"fetch/api"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/receipts/process", func(c echo.Context) error {
		errInvalidReceipt := "The receipt is invalid."
		ctx := c.Request().Context()
		req := new(api.ProcessReceiptRequest)
		if err := c.Bind(req); err != nil {
			fmt.Println("error: invalid request payload")
			return c.JSON(http.StatusBadRequest, errInvalidReceipt)
		}
		processReceiptResponse, err := ProcessReceipt(ctx, *req)
		if err != nil {
			fmt.Printf("error: %s \n", err)
			return c.JSON(http.StatusBadRequest, errInvalidReceipt)
		}
		return c.JSON(http.StatusOK, processReceiptResponse)
	})
	e.GET("/receipts/:id/points", func(c echo.Context) error {
		errNotFound := "No receipt found for that ID."
		ctx := c.Request().Context()
		id := c.Param("id")
		req := api.GetPointsRequest{ID: id}
		getPointsResponse, err := GetPoints(ctx, req)
		if err != nil {
			fmt.Printf("error: %s \n", err)
			return c.JSON(http.StatusNotFound, errNotFound)
		}
		return c.JSON(http.StatusOK, getPointsResponse)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
