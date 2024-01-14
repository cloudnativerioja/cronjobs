package main

import (
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	botToken := os.Getenv("BOT_TOKEN")
	chatIDStr := os.Getenv("CHAT_ID")

	if botToken == "" || chatIDStr == "" {
		log.Fatal("BOT_TOKEN and CHAT_ID must be set in the environment")
	}

	// Convert chatID to int64
	chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		log.Fatal("Invalid CHAT_ID. Must be a valid integer.")
	}

	// Create a new bot instance
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	// Set the chat ID to the channel ID
	chat := tgbotapi.ChatConfig{
		ChatID: chatID,
	}

	// Create a message config with Markdown formatting
	msgText := "🚀 **[DevOps Talks](https://docs.cloudnativerioja.com/recursos/assets/devops-talks.jpg) en La Rioja** 🚀\n\n"
	msgText += "Únete a nosotros en la reunión de DevOps por excelencia en La Rioja.\n\n"
	msgText += "📅 Guarda la fecha: El último miércoles de cada mes.\n\n"
	msgText += "🗣️ **Temas**:\n"
	msgText += "   - Pipelines de CI/CD\n"
	msgText += "   - Infraestructura como Código (IaC)\n"
	msgText += "   - Contenerización (Docker, Kubernetes)\n"
	msgText += "   - Observabilidad\n"
	msgText += "   - Colaboración entre equipos de Desarrollo y Operaciones\n\n"
	msgText += "📍 **Ubicación**: Virtual\n\n"
	msgText += "🕕 **Hora**: 19:30 PM - 20:30 PM\n\n"
	msgText += "👥 **Quién debería asistir**:\n"
	msgText += "   - Desarrolladores\n"
	msgText += "   - Profesionales de Operaciones\n"
	msgText += "   - ¡Cualquier persona interesada en DevOps!\n\n"
	msgText += "📢 **¡Corre la voz!** Trae a tus colegas y amigos.\n\n"
	msgText += "📝 **Regístrate**: [https://community.cncf.io/cloud-native-rioja/]\n\n"
	msgText += "📧 **Contáctanos para más información**: [cloudnativerioja@gmail.com]\n"

	// Create a message config with Markdown formatting
	msg := tgbotapi.NewMessage(chat.ChatID, msgText)
	msg.ParseMode = tgbotapi.ModeMarkdown

	// Send the message
	_, err = bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Message sent successfully!")
}
