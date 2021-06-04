package main

import (
	RandomGoodHanime "RandomGoodHanime/pkg"
	"fmt"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	Cfg := new(RandomGoodHanime.Config)
	Cfg.Port = "80"
	Cfg.Token = "1771420037:AAHrIvi1RFh5aUcID_rkS7EXOvjcCYX77sc"

	//Start webserver
	go func() {
		RandomGoodHanime.StartWebServer(Cfg.Port)
	}()

	//Bot config
	bot, err := tb.NewBot(tb.Settings{
		Token:  Cfg.Token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//Handle commands
	for i := 0; i < len(RandomGoodHanime.Categories); i++ {
		RandomGoodHanime.HandleCommands(bot, i)
	}
	bot.Handle("/start", func(m *tb.Message) {
		bot.Send(m.Sender, "Use /help to start watching good animes")
		RandomGoodHanime.Cron("/start", m.Chat.Username, m.Chat.FirstName)
	})
	bot.Handle("/help", func(m *tb.Message) {
		bot.Send(m.Sender, RandomGoodHanime.HelpMessage)
		RandomGoodHanime.Cron("/help", m.Chat.Username, m.Chat.FirstName)
	})

	fmt.Println(RandomGoodHanime.Log)

	//Start bot
	bot.Start()
}
