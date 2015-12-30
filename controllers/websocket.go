// Copyright 2013 Beego Samples authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package controllers

import (
	"encoding/json"
	// "fmt"
	"github.com/astaxie/beego"
	// "github.com/bitly/go-simplejson"
	"github.com/gorilla/websocket"
	"goAPI/models"
	"net/http"
	"strings"
)

// WebSocketController handles WebSocket requests.
type WebSocketController struct {
	beego.Controller
}

// Get method handles GET requests for WebSocketController.
func (this *WebSocketController) Get() {
	// Safe check.
	uname := this.GetString("uname")
	if len(uname) == 0 {
		return
	}
}

// Join method handles WebSocket requests for WebSocketController.
func (this *WebSocketController) Join() {
	uname := this.GetString("uname")
	if len(uname) == 0 {
		return
	}
	var upgrader = &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}
	//1, Upgrade from http request to WebSocket.
	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}

	go jionToChatRoom(uname, ws)
}

// broadcastWebSocket broadcasts messages to WebSocket users.
func broadcastWebSocket(event models.Event) {

	data, err := json.Marshal(event)
	if err != nil {
		beego.Error("Fail to marshal event:", err)
		return
	}
	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		if event.Type == models.EVENT_CHAT {
			chat := new(models.ChatBean)
			json.Unmarshal([]byte(event.Content), chat)
			beego.Info("1")
			if strings.EqualFold(chat.To, sub.Value.(Subscriber).Name) {
				beego.Info("2")
				// Immediately send event to WebSocket users.
				ws := sub.Value.(Subscriber).Conn
				if ws != nil {
					beego.Info("3")
					if ws.WriteMessage(websocket.TextMessage, data) != nil {
						// User disconnected.
						unsubscribe <- sub.Value.(Subscriber).Name
					}
				} else {
					beego.Info("4")
				}
			} else {
				beego.Info("5")
				// beego.Info(" your friend ", chat.To, "is not online")
			}
		} else {
			beego.Info("0")
			// Immediately send event to WebSocket users.
			ws := sub.Value.(Subscriber).Conn
			if ws != nil {
				if ws.WriteMessage(websocket.TextMessage, data) != nil {
					// User disconnected.
					unsubscribe <- sub.Value.(Subscriber).Name
				}
			}
		}

	}

}

func jionToChatRoom(uname string, ws *websocket.Conn) {

	// Join chat room.
	Join(uname, ws)
	defer Leave(uname)

	// Message receive loop.
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		//=======deal with chat====
		// js, err := simplejson.NewJson(p)
		// if err != nil {
		// 	panic(err.Error())
		// }
		chat := new(models.ChatBean)
		json.Unmarshal(p, chat)

		//=====================
		if chat.Type == 1 || len(chat.To) > 0 {
			publish <- newEvent(models.EVENT_CHAT, uname, string(p))
		} else {
			publish <- newEvent(models.EVENT_MESSAGE, uname, string(p))
		}

	}
}
