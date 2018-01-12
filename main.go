package main

import (
	"fmt"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	token, found := os.LookupEnv("TELEGRAM_TOKEN_BUGURTOBOT")
	if !found {
		fmt.Println("Env variables not found")
		return
	}
	bot, err := tb.NewBot(tb.Settings{
		Token: token,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Username -> Earliest time to pwn.
	cooldown := map[int]time.Time{}

	bot.Handle(tb.OnText, func(m *tb.Message) {
		lastPwn, found := cooldown[m.Contact.UserID]
		now := time.Now()
		if !found || (found && lastPwn.Before(now)) {
			fmt.Println("Replying!")
			cooldown[m.Contact.UserID] = now.Add(90 * time.Second)
			bot.Reply(m, "-")
		} else {
			fmt.Printf("Have to wait %.4f seconds...\n", lastPwn.Sub(now).Seconds())
		}
	})

	bot.Start()
}
