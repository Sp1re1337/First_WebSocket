package bot

import (
	"log"
	"bot/internal/ws"  // Імпортуємо наш WebSocket клієнт

	"github.com/mymmrac/telego"
)

// HandleMessage обробляє повідомлення користувача і надсилає його у WebSocket
func HandleMessage(bot *telego.Bot, message telego.Message, wsClient *ws.WebSocketClient) {
	log.Printf("Отримано повідомлення: %s", message.Text)
	wsClient.SendMessage(message.Text) // Надсилаємо повідомлення на WebSocket сервер
}
