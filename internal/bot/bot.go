package bot

import (
	"os"
	"os/signal"
	"strings"
	"syscall"

	embed "github.com/Clinet/discordgo-embed"
	"github.com/bwmarrin/discordgo"
	"github.com/labstack/gommon/log"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	bot  = kingpin.New("discotime", "A Discord bot to help you with timezone conversions.")
	help = kingpin.Command("help", "Show the help for this bot.")
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
	log.Infof("[%v] -> %v", m.Author.Username, m.Content)

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	cmd, err := bot.Parse(strings.Fields(m.Content))
	if err != nil {
		s.ChannelMessageSendEmbed(m.ChannelID, embed.NewErrorEmbed("Example Error", err.Error()))
	}
	switch cmd {
	case "help":
		s.ChannelMessageSendEmbed(m.ChannelID, embed.NewGenericEmbed("Usage",
			`This bot understads the following commands:
				halp
		
		`))
	}

}
