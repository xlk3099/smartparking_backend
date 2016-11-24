package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	. "github.com/xlk3099/smartparking_backend/model"
	"net/http"
)

var Parkings = map[int]Parking{
	1: {true, ""},
	2: {true, ""},
	3: {true, ""},
}

// Server : Wrap the gin.Engine type
type Server struct {
	*gin.Engine
}

// NewServer : Function to create a new gin.Default() server.
func NewServer() Server {
	server := Server{gin.Default()}
	server.LoadHTMLFiles("client_test.html")

	// For testing
	server.GET("/", func(c *gin.Context) {
		c.HTML(200, "client_test.html", nil)
	})
	// Restful api
	server.GET("/parking", getParkings)
	server.PUT("/parking/:id", updateParking)

	// Websocket
	server.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})

	return server
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func getParkings(c *gin.Context) {

}

func updateParking(c *gin.Context) {
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, msg)
	}
}
