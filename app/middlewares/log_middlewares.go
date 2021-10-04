package middlewares

import (
	"backend/drivers/mongodb"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func Log(c echo.Context, reqBody, resBody []byte) {
	ctx := c.Request().Context()
	db := mongodb.Connect(ctx)

	RequestURI := c.Request().RequestURI
	RequestMethod := c.Request().Method

	//mongoDB connect disini diambil dari global

	db.Collection(viper.GetString(`mongodb.collection`)).InsertOne(ctx, mongodb.LogData{RequestURI, RequestMethod})
}
