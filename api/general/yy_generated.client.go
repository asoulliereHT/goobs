// This file has been automatically generated. Don't edit it.
//go:generate go run github.com/andreykaipov/funcopgen -type Client -prefix With -factory -unexported

package general

import websocket "github.com/gorilla/websocket"

// Client represents a client for 'general' requests
type Client struct {
	conn *websocket.Conn
}
