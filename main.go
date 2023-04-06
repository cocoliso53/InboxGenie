package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func checkEmails(bot *tgbotapi.BotAPI, tgAPI string, server string,
	email string, password string, tgID int64) {

	imapClient, err := client.DialTLS(server, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer imapClient.Logout()

	if err := imapClient.Login(email, password); err != nil {
		log.Fatal(err)
	}

	_, err = imapClient.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}

	criteria := imap.NewSearchCriteria()
	criteria.WithoutFlags = []string{imap.SeenFlag}
	unseenSeqNums, err := imapClient.Search(criteria)
	if err != nil {
		log.Fatal(err)
	}

	if len(unseenSeqNums) > 0 {

		tgMessage := tgbotapi.NewMessage(tgID, "You got mail!")
		_, err = bot.Send(tgMessage)
		if err != nil {
			log.Println("Error sending Telegram message:", err)
		}
	}
}

func main() {

	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading config")
	}

	server := os.Getenv("SERVER")
	email := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")
	tgIDString := os.Getenv("TGID")
	tgAPI := os.Getenv("TGAPI")

	tgID, err := strconv.ParseInt(tgIDString, 10, 64)
	if err != nil {
		log.Println("Error:", err)
	}

	bot, err := tgbotapi.NewBotAPI(tgAPI)
	if err != nil {
		log.Panic(err)
	}

	emailCheckInterval := 1 * time.Minute
	ticker := time.NewTicker(emailCheckInterval)

	go func() {
		for range ticker.C {
			checkEmails(bot, tgAPI, server, email, password, tgID)
		}
	}()

	select {}
}
