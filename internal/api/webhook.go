package api

import (
	"encoding/json"
	"log"
	"time"

	"github.com/BrunoKrugel/go-webhook/internal/client"
	"github.com/BrunoKrugel/go-webhook/internal/model"
	"github.com/labstack/echo"
)

func Webhook(c echo.Context) error {

	log.Println(c.Request().Header)

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

	kiwifyLog, _ := json.Marshal(kiwify)
	log.Println(string(kiwifyLog))

	date, _ := time.Parse("2006-01-02", kiwify.CreatedAt[0:10])

	kiwifySales := model.MongoSales{
		UserId:      c.Param("user"),
		ProductId:   kiwify.Product.ProductID,
		StoreId:     kiwify.StoreID,
		Date:        date,
		ProductName: kiwify.Product.ProductName,
	}

	kiwifySalesInfo := model.MongoSalesInfo{
		UserId:        c.Param("user"),
		ProductName:   kiwify.Product.ProductName,
		StoreId:       kiwify.StoreID,
		Date:          date,
		PaymentMethod: kiwify.PaymentMethod,
		OrderStatus:   kiwify.OrderStatus,
		Commission:    kiwify.Commissions.MyCommission,
	}

	go func() {
		client.UpdateSales(kiwifySales)
	}()

	go func() {
		client.InsertSales(kiwifySalesInfo)
	}()
	return c.JSON(200, "Ok")
}

func Web(c echo.Context) error {
	return c.JSON(404, "Not authorized")
}
