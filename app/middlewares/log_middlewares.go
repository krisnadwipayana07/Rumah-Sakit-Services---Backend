package middlewares

import (
	"backend/drivers/mongodb"
	"context"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func Log(c echo.Context, reqBody, resBody []byte) {
	var ctx context.Context
	db, err := mongodb.Connect(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	RequestURI := c.Request().RequestURI
	RequestMethod := c.Request().Method

	//mongoDB connect disini diambil dari global

	_, err = db.Collection(viper.GetString(`mongodb.collection`)).InsertOne(ctx, mongodb.LogData{RequestURI, RequestMethod})
	if err != nil {
		log.Fatal(err.Error())
	}
}
