package main

import (
	"fmt"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	token, found := os.LookupEnv("BUGURTOBOT_TOKEN")
	if !found {
		fmt.Println("Env variables not found")
		return
	}
	bot, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Username -> Earliest time to pwn.
	permittedUsers := map[string]time.Time{
		"yabalaban":  time.Now(),
		"massaraksh": time.Now(),
	}

	bot.Handle(tb.OnText, func(m *tb.Message) {
		fmt.Printf("Received a message %s from %s\n", m.Text, m.Sender.Username)
		lastPwnTime, found := permittedUsers[m.Sender.Username]
		if found {
			now := time.Now()
			if lastPwnTime.Before(now) {
				fmt.Println("Replying!")
				permittedUsers[m.Sender.Username] = now.Add(95 * time.Second)
				bot.Reply(m, "-")
			} else {
				fmt.Printf("Have to wait %.4f seconds...\n", lastPwnTime.Sub(now).Seconds())
			}
		} else {
			fmt.Println("User not found")
		}
	})

	bot.Start()
}
