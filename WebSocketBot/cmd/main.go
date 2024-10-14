package main

import (
	"log"
	"os"

	"bot/internal/bot"
	"bot/internal/ws"

	"github.com/mymmrac/telego"
)

func main() {
	// Отримання токена бота з оточення
	botToken := os.Getenv("TELEGRAM_TOKEN")
	tBot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		log.Fatal(err)
	}

	// Ініціалізація WebSocket клієнта
	wsURL := "wsToken" // Заміни на реальний WebSocket URL
	wsClient := ws.NewWSClient(wsURL)
	go wsClient.Connect()

	// Запуск отримання оновлень (повідомлень) від Telegram
	updates, _ := tBot.UpdatesViaLongPolling(nil)

	// Обробка кожного оновлення
	for update := range updates {
		if update.Message != nil {
			// Викликаємо обробник повідомлень
			bot.HandleMessage(tBot, *update.Message, wsClient)
		}
	}

	// Очищення ресурсів
	tBot.StopLongPolling()
}
