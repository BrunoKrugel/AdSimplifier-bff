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
		log.Println("Empty user")
		return c.JSON(404, "Not ok")
	}

	kiwify := &model.KiwifyRequest{}
	if err := c.Bind(kiwify); err != nil {
		log.Println("Error binding" + err.Error())
		return c.JSON(200, err)
	}

	if (kiwify == &model.KiwifyRequest{}) {
		log.Println("Empty request")
		return c.JSON(404, "Not ok")
	}

	date, _ := time.Parse("2006-01-02", kiwify.CreatedAt[0:10])

	kiwifySales := model.MongoRequest{
		UserId:      c.Param("user"),
		ProductId:   kiwify.Product.ProductID,
		StoreId:     kiwify.StoreID,
		Date:        date,
		ProductName: kiwify.Product.ProductName,
	}

	client.UpdateSales(kiwifySales)

	return c.JSON(200, "Ok")
}

func Web(c echo.Context) error {
	return c.JSON(200, "Ok")
}
