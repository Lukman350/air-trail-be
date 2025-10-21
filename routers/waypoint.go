package routers

import (
	"air-trail-backend/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

var waypointRoute = Router{
	Name:     "Waypoint Endpoint",
	Endpoint: "/waypoints",
	Handler:  waypointHandler,
	Method:   GET,
}

func init() {
	ROUTERS = append(ROUTERS, waypointRoute)
}

func waypointHandler(ctx *gin.Context) {
	waypoints, err := api.Waypoint_Get()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, waypoints)
}
