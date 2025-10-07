package routers

import (
	"air-trail-backend/api"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var JetPhotosRouter Router = Router{
	Name:     "JetPhotos API",
	Endpoint: "/jet_photos",
	Handler:  jetPhotosHandler,
	Method:   GET,
}

type JetPhotosQueryParams struct {
	Registration string `form:"reg"`
}

func init() {
	ROUTERS = append(ROUTERS, JetPhotosRouter)
}

func jetPhotosHandler(ctx *gin.Context) {
	var params JetPhotosQueryParams

	err := ctx.ShouldBind(&params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, err := api.GetJetPhotos(params.Registration)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "There was an error when get jet photos"})
		log.Printf("%s", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, result)
}
