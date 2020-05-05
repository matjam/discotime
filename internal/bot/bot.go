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

const format = "3:04pm on Monday, 02 January 2006 (UTC-0700)"

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	var command string
	var reply *discordgo.MessageEmbed

	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		log.Errorf("unable to get channel: %v", err.Error())
	}

	// We only handle messages from channels and DMs.
	if channel.Type != discordgo.ChannelTypeDM && channel.Type != discordgo.ChannelTypeGuildText {
		return
	}

	// ignore messagesif they come from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// if a message was received on a channel without a ! we ignore it
	if !strings.HasPrefix(m.Content, "!") && channel.Type == discordgo.ChannelTypeGuildText {
		return
	}

	log.Infof("[%v] [%v]: %v", m.ChannelID, m.Author.Username, m.Content)

	args := strings.Fields(m.Content)
	if len(args) < 1 {
		return
	}

	// strip the bang
	if strings.HasPrefix(m.Content, "!") {
		command = strings.ToLower(args[0])[1:]
	} else {
		command = strings.ToLower(args[0])
	}

	switch command {
	case "help":
		reply = notImplementedEmbed
	case "time":
		var b strings.Builder

		now := time.Now()
		local, _ := time.LoadLocation("America/Los_Angeles")
		fmt.Fprintf(&b, "The current time is %v\n", now.Format(format))
		fmt.Fprintf(&b, "Local time is %v", now.In(local).Format(format))

		s.ChannelMessageSend(m.ChannelID, b.String())
	case "localtime":
		reply = notImplementedEmbed
	case "set":
		reply = notImplementedEmbed
	case "convert":
		reply = notImplementedEmbed
	case "remindme":
		reply = notImplementedEmbed
	}

	// only send the message if we have one to send
	if reply != nil {
		s.ChannelMessageSendEmbed(m.ChannelID, reply)
	}

	// We ignore messages that aren't proper commands as they may be for another bot
}

func setTimezone() {

}
