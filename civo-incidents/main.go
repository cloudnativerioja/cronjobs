package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mmcdole/gofeed"
)

func main() {
	telegramToken := os.Getenv("BOT_TOKEN")
	chatIDStr := os.Getenv("CHAT_ID_ADMIN")
	rssURL := "https://status.civo.com/index.xml"
	// Initialize Telegram bot
	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Fatal(err)
	}
	// Convert chatID to int64
	chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		log.Fatal("Invalid CHAT_ID. Must be a valid integer.")
	}

	// Set up RSS feed parser
	fp := gofeed.NewParser()

	// Fetch and parse the RSS feed
	feed, err := fp.ParseURL(rssURL)
	if err != nil {
		log.Printf("Error parsing RSS feed: %v\n", err)
	} else {
		// Get the current hour
		currentHour := time.Now().Hour()

		// Check for incidents in the feed within the current hour
		for _, item := range feed.Items {
			if strings.Contains(strings.ToLower(item.Title), "[resolved]") {
				// Skip resolved incidents
				continue
			}

			// Check if the incident occurred within the current hour
			if item.PublishedParsed != nil && item.PublishedParsed.Hour() == currentHour {
				// Send notification to Telegram channel for each incident
				message := fmt.Sprintf("🚨 Incident Detected: %s\n%s", item.Title, item.Link)
				sendTelegramMessage(bot, chatID, message)
			}
		}
	}
}

func sendTelegramMessage(bot *tgbotapi.BotAPI, chatID int64, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message to Telegram: %v\n", err)
	}
}
