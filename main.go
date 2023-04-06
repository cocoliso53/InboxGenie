package main

import (
	"log"
	"time"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func checkEmails(bot *tgbotapi.BotAPI) {

	imapClient, err := client.DialTLS("", nil)
	if err != nil {
		log.Fatal(err)
	}

	defer imapClient.Logout()

	if err := imapClient.Login("", ""); err != nil {
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
		var tgUserID int64
		tgUserID = 
		tgMessage := tgbotapi.NewMessage(tgUserID,"You got mail!")
			_, err := bot.Send(tgMessage)
			if err != nil {
				log.Println("Error sending Telegram message:", err)
			}
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI("")
	if err != nil {
		log.Panic(err)
	}

	emailCheckInterval := 1 * time.Minute
	ticker := time.NewTicker(emailCheckInterval)

	go func() {
		for range ticker.C {
			checkEmails(bot)
		}
	}()

	select {}
}
