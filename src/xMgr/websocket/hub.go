// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import (
	"time"
	"xsw/go_pub/x"
)

type MsgBuf struct {
	ID      int
	ByteMsg []byte
}

// hub maintains the set of active connections and broadcasts messages to the
// connections.
type hub struct {
	// Registered connections.
	connections    map[*connection]bool
	connections_id map[int]*connection

	// Inbound messages from the connections.
	broadcast chan []byte

	// Register requests from the connections.
	register chan *connection

	// Unregister requests from connections.
	unregister chan *connection

	chanSend chan *MsgBuf

	chanRcv chan *MsgBuf
}

var h = hub{
	broadcast:      make(chan []byte, 512),
	register:       make(chan *connection, 512),
	unregister:     make(chan *connection, 512),
	chanSend:       make(chan *MsgBuf, 512),
	chanRcv:        make(chan *MsgBuf, 512),
	connections:    make(map[*connection]bool, 512),
	connections_id: make(map[int]*connection, 512),
}

func PeekRcvMsg() (*MsgBuf, bool) {
	select {
	case msg := <-h.chanRcv:
		if msg == nil {
			x.PrintErr("msg is nil")
			return nil, true
		}
		return msg, false
		break
	//超时，避免过忙
	case <-time.After(100 * time.Millisecond):
		break
	}
	return nil, false
}

func SendMsg(cid int, byteMsg []byte) *x.Error {
	if byteMsg == nil {
		return x.XErrStr("byteMsg is nil")
	}
	h.chanSend <- &MsgBuf{ByteMsg: byteMsg, ID: cid}
	return nil
}

func send_connection(msg *MsgBuf) *x.Error {
	if msg == nil {
		return x.XErrStr("msg is nil")
	}
	c := h.connections_id[msg.ID]
	if c == nil {
		return x.XErrStr("connection is nil")
	}
	c.send <- msg.ByteMsg
	return nil
}

func add_connection(c *connection) {
	h.connections[c] = true
	h.connections_id[c.nID] = c
}

func del_connection(nID int) {
	c := h.connections_id[nID]
	if c != nil {
		c.close()
		delete(h.connections, c)
	}
	delete(h.connections_id, nID)
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			add_connection(c)

		case c := <-h.unregister:
			del_connection(c.nID)

		case m := <-h.broadcast:
			for c := range h.connections {
				select {
				case c.send <- m:
				default:
					del_connection(c.nID)
				}
			}

		case msg := <-h.chanSend:
			send_connection(msg)
		}
	}
}
