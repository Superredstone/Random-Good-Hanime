package RandomGoodHanime

import (
	"strings"

	tb "gopkg.in/tucnak/telebot.v2"
)

func HandleCommands(b *tb.Bot, i int) {
	b.Handle(Categories[i], func(m *tb.Message) {
		CategoriesWithoutSlash := strings.Replace(Categories[i], "/", "", 1)

		photo := &tb.Photo{File: tb.FromURL(RetrieveHentai(CategoriesWithoutSlash))}

		b.Send(m.Sender, photo)
		Cron(Categories[i], m.Chat.Username, m.Chat.FirstName)
	})
}
