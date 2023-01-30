package api

import (
	"fmt"

	"github.com/BrunoKrugel/go-webhook/internal/client"
	"github.com/BrunoKrugel/go-webhook/internal/model"
	"github.com/labstack/echo"
)

func Webhook(c echo.Context) error {
	user := c.Param("user")
	kiwify := &model.KiwifyRequest{}
	if err := c.Bind(kiwify); err != nil {
		return err
	}

	kiwifySales := model.MongoRequest{
		UserId:      user,
		ProductId:   kiwify.Product.ProductID,
		StoreId:     kiwify.StoreID,
		Date:        kiwify.CreatedAt[0:10],
		ProductName: kiwify.Product.ProductName,
	}

	client.UpdateSales(kiwifySales)

	fmt.Printf("User: %s", user)
	return c.JSON(200, "Hello, World!")
}
