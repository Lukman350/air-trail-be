package routers

import (
	"air-trail-backend/database/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var AirportRoute Router = Router{
	Name:     "Airport Route",
	Endpoint: "/airports",
	Method:   GET,
	Handler:  AirportRouteHandler,
}

func init() {
	ROUTERS = append(ROUTERS, AirportRoute)
}

func AirportRouteHandler(ctx *gin.Context) {
	result, err := models.Airport_GetAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
