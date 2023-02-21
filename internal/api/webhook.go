package api

import (
	"fmt"
	"log"
	"time"

	"github.com/BrunoKrugel/go-webhook/internal/client"
	"github.com/BrunoKrugel/go-webhook/internal/model"
	"github.com/labstack/echo"
)

func Webhook(c echo.Context) error {
	log.Println("Webhook called at: " + time.Now().String())
	kiwify := &model.KiwifyRequest{}
	if err := c.Bind(kiwify); err != nil {
		fmt.Println(err)
		return c.JSON(200, err)
	}

	if (kiwify == &model.KiwifyRequest{}) {
		return c.JSON(200, "Not ok")
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
