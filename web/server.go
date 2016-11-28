package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	. "github.com/xlk3099/smartparking_backend/model"
)

var Parkings = map[int]Parking{
	1: {1, true, "IPHONE6"},
	2: {2, true, "SG50"},
	3: {3, true, "ABC123"},
}

// Server : Wrap the gin.Engine type
type Server struct {
	*gin.Engine
}

// NewServer : Function to create a new gin.Default() server.
func NewServer() Server {
	server := Server{gin.Default()}
	server.LoadHTMLFiles("client.html")

	// For testing
	server.GET("/", func(c *gin.Context) {
		c.HTML(200, "client.html", nil)
	})
	// Restful api
	server.GET("/parking", getParkings)
	server.PUT("/parking/:id", updateParking)

	// Websocket
	// server.GET("/ws", func(c *gin.Context) {
	// 	wshandler(c.Writer, c.Request)
	// })
	server.GET("/ws", wshandler)

	return server
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func getParkings(c *gin.Context) {
	c.JSON(http.StatusOK, Parkings)
}

func updateParking(c *gin.Context) {
}

func wshandler(c *gin.Context) {
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
		if err = conn.WriteJSON(Parkings); err != nil {
			fmt.Println(err)
		}
	}
}
