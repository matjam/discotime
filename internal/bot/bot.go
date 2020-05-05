package bot

import (
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

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.GuildID != "" || m.Author.ID == s.State.User.ID {
		return
	}
	log.Infof("[%v] -> %v", m.Author.Username, m.Content)

	var e *discordgo.MessageEmbed

	args := strings.Fields(m.Content)
	if len(args) < 1 {
		return
	}
	switch args[0] {
	case "help":
		e = embed.NewGenericEmbed("Usage", "Supported commands:\n"+
			" * time now - get the current time in UTC")
	case "time":
		now := time.Now()
		e = embed.NewGenericEmbed("Time now UTC:", now.Format(time.RFC1123Z))
	default:
		e = embed.NewErrorEmbed("Huh?", "Sorry, I don't understand that.")
	}

	s.ChannelMessageSendEmbed(m.ChannelID, e)

}
