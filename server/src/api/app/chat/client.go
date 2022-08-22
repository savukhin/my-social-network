package chat

import (
	"api/app/utils"
	"api/dto"
	"api/mappers"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

type Client struct {
	ws      *websocket.Conn
	hub     *Hub
	chat_id int
	send    chan *dto.Message
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (client Client) Write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		err := client.ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	for {
		select {
		case message, ok := <-client.send:
			if !ok {
				err := client.ws.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					fmt.Println(err)
				}
				return
			}

			err := client.ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				panic(err)
			}

			writer, err := client.ws.NextWriter(websocket.TextMessage)
			if err != nil {
				writer.Close()
				return
			}

			b, err := json.Marshal(message)
			if err != nil {
				writer.Close()
				continue
			}

			_, err = writer.Write(b)
			if err != nil {
				writer.Close()
				continue
			}

			if err := writer.Close(); err != nil {
				return
			}
		case <-ticker.C:
			err := client.ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				log.Printf("error: %v", err)
			}
			if err := client.ws.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (client *Client) Read() {
	defer func() {
		client.hub.unregister <- client
		err := client.ws.Close()
		if err != nil {
			fmt.Println("Error closing client ", err)
		}
	}()

	err := client.ws.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		fmt.Println("Error setting read deadline on client ", err)
	}

	client.ws.SetPongHandler(
		func(string) error {
			client.ws.SetReadDeadline(time.Now().Add(pongWait))
			return nil
		},
	)

	for {
		_, message_b, err := client.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Println("Error reading message: ", err)
			}
			break
		}

		message_b = bytes.TrimSpace(bytes.Replace(message_b, []byte{'\n'}, []byte{' '}, -1))

		message_input := &dto.MessageInput{}
		if err = json.Unmarshal(message_b, &message_input); err != nil {
			fmt.Println(err)
			break
		}

		_, err = utils.UnpackJWT(message_input.Token)
		if err != nil {
			fmt.Println(err)
			break
		}

		content, err := mappers.MessageToContent(message_input)
		if err != nil {
			fmt.Println(err)
			continue
		}
		content.Save()

		message, _ := mappers.ToMessage(content)

		client.hub.send <- message
	}
}

func ServeWs(hub *Hub, res http.ResponseWriter, req *http.Request) {
	chat_id, err := strconv.Atoi(mux.Vars(req)["chat_id"])
	if err != nil {
		utils.ResponseError(res, err, http.StatusBadGateway)
		return
	}

	conn, err := upgrader.Upgrade(res, req, nil)

	if err != nil {
		utils.ResponseError(res, err, http.StatusBadGateway)
		return
	}

	client := &Client{
		ws:      conn,
		hub:     hub,
		chat_id: chat_id,
		send:    make(chan *dto.Message),
	}

	hub.register <- client

	go client.Write()
	go client.Read()
}
