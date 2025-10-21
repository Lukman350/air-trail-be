package routers

import (
	"air-trail-backend/api"
	"air-trail-backend/utils"
	"encoding/json"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

var Cat021Router Router = Router{
	Name:     "Cat021 Track",
	Endpoint: "/ws/cat021-track",
	Handler:  cat021Handler,
	Method:   GET,
}

func init() {
	ROUTERS = append(ROUTERS, Cat021Router)
}

func cat021Handler(ctx *gin.Context) {

	cat021Ws := WebSocket{
		Name:          "Cat021 Service",
		OnReadMessage: readWsMessage,
	}

	if err := cat021Ws.Connect(ctx.Writer, ctx.Request, nil); err != nil {
		log.Println("WebSocket connect error:", err)
		return
	}

	defer cat021Ws.Disconnect()

	go cat021Ws.ReadLoop()

	cat021 := api.Cat021{}

	go sendWsMessage(&cat021Ws, api.Cat021Channel)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("Client disconnected")
			cat021Ws.Disconnect()
			return
		case <-ticker.C:
			cat021.Get()
		}
	}

}

func readWsMessage(mt int, msg []byte, err error, ws *WebSocket) {
	if err != nil {
		switch err.(type) {
		case *utils.Cat021Error:
			log.Println(err.Error())
		}
		return
	}

	// log.Printf("Received message (mt=%d): %s\n", mt, msg)

	var bbox utils.BBox
	if err := json.Unmarshal(msg, &bbox); err == nil {
		ws.BBox = &bbox
		return
	}
}

func sendWsMessage(ws *WebSocket, ch <-chan api.Cat021) {
	for data := range ch {
		// skip if outside bbox
		// log.Printf("bbox: %+v\n", ws.BBox)
		if ws.BBox != nil && !ws.BBox.Contains(data.Latitude, data.Longitude) {
			continue
		}

		prevRaw, ok := api.Cat021Cache.Load(data.IcaoAddress)
		if ok {
			prev := prevRaw.(api.Cat021)
			prevTs, _ := time.Parse(time.RFC3339, prev.UpdateTimestamp.Format(time.RFC3339))
			newTs, _ := time.Parse(time.RFC3339, data.UpdateTimestamp.Format(time.RFC3339))

			if !newTs.After(prevTs) {
				continue
			}
		}

		if *data.UpdateDelete == "DELETE" {
			api.Cat021Cache.Delete(data.IcaoAddress)

			if err := ws.SendMessage(map[string]any{
				"icaoAddress": data.IcaoAddress,
				"delete":      true,
			}); err != nil {
				log.Println(err.Error())
			}

			continue
		}

		api.Cat021Cache.Store(data.IcaoAddress, data)

		if err := ws.SendMessage(&data); err != nil {
			switch err.(type) {
			case *utils.Cat021Error:
				log.Println(err.Error())
			}
			return
		}
	}
}
