package api

import (
	"github.com/BrunoKrugel/go-webhook/internal/client"
	"github.com/BrunoKrugel/go-webhook/internal/model"
	"github.com/labstack/echo"
)

func Webhook(c echo.Context) error {

	kiwify := &model.KiwifyRequest{}
	if err := c.Bind(kiwify); err != nil {
		return err
	}

	if (kiwify == &model.KiwifyRequest{}) {
		return c.JSON(200, "Not ok")
	}

	kiwifySales := model.MongoRequest{
		UserId:      c.Param("user"),
		ProductId:   kiwify.Product.ProductID,
		StoreId:     kiwify.StoreID,
		Date:        kiwify.CreatedAt[0:10],
		ProductName: kiwify.Product.ProductName,
	}

	client.UpdateSales(kiwifySales)

	return c.JSON(200, "Ok")
}
