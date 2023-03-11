package String

import (
	"encoding/json"
	"log"

	"github.com/labstack/echo"
)

func GetJSONRawBody(c echo.Context) map[string]interface{} {

	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		log.Println("Error decoding request body" + err.Error())
		return nil
	}

	return jsonBody
}
