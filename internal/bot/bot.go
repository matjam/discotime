package bot

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	embed "github.com/Clinet/discordgo-embed"
	"github.com/bwmarrin/discordgo"
	"github.com/labstack/gommon/log"
)

// Run will start the discord bot
func Run(token string) {
	log.Info("starting discord bot")
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Errorf("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Errorf("error opening connection,", err)
		return
	}

	log.Info("Bot is now running.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()

}

var notImplementedEmbed = embed.NewGenericEmbed("Not Implemented", "```\nnot yet implemented.\n```")

const format = "3:04pm on Monday, 02 January 2006 MST (UTC-0700)"

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// ignore messages not prefixed with ! or if they come from the bot itself
	if !strings.HasPrefix(m.Content, "!") || m.Author.ID == s.State.User.ID {
		return
	}
	log.Infof("[%v] -> %v", m.Author.Username, m.Content)

	var e *discordgo.MessageEmbed

	args := strings.Fields(m.Content)
	if len(args) < 1 {
		return
	}

	command := strings.ToLower(args[0])[1:]

	switch command {
	case "help":
		e = notImplementedEmbed
	case "time":
		now := time.Now()
		e = embed.NewGenericEmbed("Current Time UTC", fmt.Sprintf("```\n%v\n```", now.Format(format)))
	case "localtime":
		e = notImplementedEmbed
	case "set":
		e = notImplementedEmbed
	case "convert":
		e = notImplementedEmbed
	case "remindme":
		e = notImplementedEmbed
	}

	// only send the message if we have one to send
	if e != nil {
		s.ChannelMessageSendEmbed(m.ChannelID, e)
	}

	// We ignore messages that aren't proper commands as they may be for another bot
}
