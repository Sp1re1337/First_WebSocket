package ws

import (
	"log"
	"github.com/gorilla/websocket"
)

// WebSocketClient представляє клієнта для WebSocket
type WebSocketClient struct {
	url         string
	conn        *websocket.Conn
	messageChan chan string
}

// NewWSClient створює нового WebSocket клієнта
func NewWSClient(url string) *WebSocketClient {
	return &WebSocketClient{
		url:         url,
		messageChan: make(chan string),
	}
}

// Connect підключає клієнта до WebSocket сервера
func (c *WebSocketClient) Connect() {
	var err error
	c.conn, _, err = websocket.DefaultDialer.Dial(c.url, nil)
	if err != nil {
		log.Fatal("Помилка підключення до WebSocket:", err)
	}
	log.Println("Підключено до WebSocket")

	go func() {
		for {
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				log.Println("Помилка читання повідомлення:", err)
				return
			}
			c.messageChan <- string(message)
		}
	}()
}

// SendMessage надсилає повідомлення на сервер через WebSocket
func (c *WebSocketClient) SendMessage(msg string) {
	err := c.conn.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		log.Println("Помилка надсилання повідомлення:", err)
	}
}

// OnMessage обробляє вхідні повідомлення з WebSocket
func (c *WebSocketClient) OnMessage(handler func(string)) {
	for msg := range c.messageChan {
		handler(msg)
	}
}
