package routers

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	PATCH  Method = "PATCH"
	DELETE Method = "DELETE"
)

var ROUTERS []Router

type IRouter interface {
	GetEndpoint() string
	GetHandler() func(*gin.Context)
}

type Router struct {
	Name     string
	Endpoint string
	Handler  func(*gin.Context)
	Method
}

func InitRouters(router *gin.Engine) {
	for _, route := range ROUTERS {
		log.Printf("Creating route for %s at %s", route.Name, route.Endpoint)

		switch method := route.Method; method {
		case GET:
			router.GET(route.Endpoint, route.Handler)
		case POST:
			router.POST(route.Endpoint, route.Handler)
		case PATCH:
			router.PATCH(route.Endpoint, route.Handler)
		case PUT:
			router.PUT(route.Endpoint, route.Handler)
		case DELETE:
			router.DELETE(route.Endpoint, route.Handler)
		}
	}
}
