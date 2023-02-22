package api

import (
	"log"
	"time"

	"github.com/BrunoKrugel/go-webhook/internal/client"
	"github.com/BrunoKrugel/go-webhook/internal/model"
	"github.com/labstack/echo"
)

func Webhook(c echo.Context) error {
	log.Println("Webhook called by user: " + c.Param("user"))

	if c.Param("user") == "" {
		log.Println("Empty user received")
		return c.JSON(404, "Empty user received")
	}

	if c.Param("user") == "favicon.ico" {
		log.Println("Web request received.")
		return c.JSON(404, "Web request received.")
	}

	kiwify := &model.KiwifyRequest{}
	if err := c.Bind(kiwify); err != nil {
		log.Println("Error binding" + err.Error())
		return c.JSON(200, err)
	}

	if kiwify.Product.ProductID == "" {
		log.Println("Empty request received.")
		return c.JSON(404, "Empty request received.")
	}

	date, _ := time.Parse("2006-01-02", kiwify.CreatedAt[0:10])

	kiwifySales := model.MongoRequest{
		UserId:      c.Param("user"),
		ProductId:   kiwify.Product.ProductID,
		StoreId:     kiwify.StoreID,
		Date:        date,
		ProductName: kiwify.Product.ProductName,
	}

	go func() {
		client.UpdateSales(kiwifySales)
	}()

	return c.JSON(200, "Ok")
}

func Web(c echo.Context) error {
	return c.JSON(404, "Not authorized")
}
