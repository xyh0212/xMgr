// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
	"xsw/go_pub/x"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 5120
)

var upgrader = websocket.Upgrader{
	CheckOrigin:     func(r *http.Request) bool { return true },
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var (
	g_nConnectionPawnID int = 0
)

// connection is an middleman between the websocket connection and the hub.
type connection struct {
	nID int
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

func (c *connection) close() {
	close(c.send)
}

// readLoop pumps messages from the websocket connection to the hub.
func (c *connection) readLoop() {
	defer func() {
		h.unregister <- c
		c.ws.Close()
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, b, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
		msg := MsgBuf{
			ID:      c.nID,
			ByteMsg: b,
		}
		h.chanRcv <- &msg
		x.PrintInfo(c.nID, " rcv: ", string(b))
	}

	x.PrintInfo(c.nID, " disconnected")
}

// write writes a message with the given message type and payload.
func (c *connection) write(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	e := c.ws.WriteMessage(mt, payload)
	x.PrintDbg(c.nID, "send: ", string(payload))
	return e
}

// writeLoop pumps messages from the hub to the websocket connection.
func (c *connection) writeLoop() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func serveWs(w http.ResponseWriter, r *http.Request) {
	g_nConnectionPawnID += 1
	x.PrintInfo(g_nConnectionPawnID, ",", r.RemoteAddr)

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		x.PrintErr(g_nConnectionPawnID, ",", err.Error())
		return
	}

	c := &connection{nID: g_nConnectionPawnID, send: make(chan []byte, 2048), ws: ws}
	h.register <- c
	go c.writeLoop()
	c.readLoop()
}
