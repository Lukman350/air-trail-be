package routers

import (
	"air-trail-backend/api"
	"air-trail-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var AftnRouter Router = Router{
	Name:     "AFTN Data",
	Endpoint: "/:callsign/aftn",
	Method:   GET,
	Handler:  AftnHandler,
}

func init() {
	ROUTERS = append(ROUTERS, AftnRouter)
}

type CallsignBinding struct {
	Callsign string `uri:"callsign" json:"callsign" binding:"required"`
}

func AftnHandler(ctx *gin.Context) {
	var callsign CallsignBinding

	if err := ctx.ShouldBindUri(&callsign); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var result = api.Aftn{}

	if err := result.GetByCallsign(callsign.Callsign); err != nil {
		switch err.(type) {
		case *utils.NotFoundError:
			ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, result)
}
