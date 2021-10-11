package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/bwmarrin/discordgo"
)

var BotID string

func main() {
	dg, err := discordgo.New("Bot " + os.Args[1])

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := dg.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotID = u.ID

	dg.AddHandler(messageHandler)

	err = dg.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Running")
	<-make(chan struct{})
	return
}

func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == BotID {
		return
	}
	if message.Content == "!spurgt" {
		_, _ = session.ChannelMessageSend(message.ChannelID, "Hvem spurgte dig?")
	}
	if message.Author.ID == "247821133758464002" {
		var num = rand.Intn(100)
		if num > 95 {
			_, _ = session.ChannelMessageSend(message.ChannelID, "Fandme autistisk sagt")
		}
	}
	fmt.Println(message.Content)
}
