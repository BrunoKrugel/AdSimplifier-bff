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
		User_id:      c.Param("user"),
		Product_id:   kiwify.Product.ProductID,
		Store_id:     kiwify.StoreID,
		Date:         date,
		Product_name: kiwify.Product.ProductName,
	}

	kiwifySalesInfo := model.MongoSalesInfo{
		User_id:        c.Param("user"),
		Product_name:   kiwify.Product.ProductName,
		Store_id:       kiwify.StoreID,
		Date:           date,
		Payment_method: kiwify.PaymentMethod,
		Order_status:   kiwify.OrderStatus,
		Commission:     kiwify.Commissions.MyCommission,
		Order_ref:      kiwify.OrderRef,
		Src:            kiwify.TrackingParameters.Src,
		Sck:            kiwify.TrackingParameters.Sck,
		Utm_medium:     kiwify.TrackingParameters.UtmMedium,
		Utm_content:    kiwify.TrackingParameters.UtmContent,
		Utm_campaign:   kiwify.TrackingParameters.UtmCampaign,
	}

	go func() {
		client.UpdateSales(kiwifySales)
	}()

	go func() {
		client.InsertSales(kiwifySalesInfo)
	}()
	return c.JSON(200, "Ok")
}
