package bot

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/matjam/discotime/internal/cache"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Run will start the discord bot
func Run(token string) {
	log.Info().Msg("starting discord bot")
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Error().Msgf("error creating Discord session: %v", err.Error())
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Error().Msgf("error opening connection: %v", err.Error())
		return
	}

	log.Info().Msg("Bot is now running.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()

}

const format = "**3:04pm** on **Monday, 02 January 2006** (UTC-0700)"

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	var command string

	userID := fmt.Sprintf("%v#%v", m.Author.Username, m.Author.Discriminator)

	sublogger := log.With().Str("channel_id", m.ChannelID).Str("user_id", userID).Str("content", m.Content).Logger()

	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		sublogger.Error().Msgf("unable to get channel: %v", err.Error())
	}

	user, err := s.User(m.Author.ID)
	if err != nil {
		sublogger.Error().Msgf("unable to get user: %v", err.Error())

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

	sublogger.Info().Msgf("processing command")
	ctx := discordContext{
		session:   s,
		user:      user,
		userID:    userID,
		channelID: m.ChannelID,
		logCtx:    &sublogger,
	}

	switch command {
	case "help":
		ctx.reply("")
	case "time":
		ctx.getTime()
	case "localtime":
		//
	case "set":
		ctx.setTimezone(args[1:])
	case "get":
		ctx.show()
	case "convert":
		//
	case "remindme":
		//
	}

	// We ignore messages that aren't proper commands as they may be for another bot
}

type discordContext struct {
	session   *discordgo.Session
	user      *discordgo.User
	userID    string
	channelID string
	logCtx    *zerolog.Logger
}

func (ctx *discordContext) setTimezone(args []string) {
	if len(args) == 0 {
		ctx.reply("`set` requires you to provide a timezone argument.")
		return
	}
	ctx.log().Info().Msgf("handling set request: %v", args)
	location, err := time.LoadLocation(args[0])
	if err != nil {
		ctx.log().Error().Msgf("couldn't parse timezone [%v]: %v", args[0], err.Error())
		ctx.reply(fmt.Sprintf("Sorry, *%v* is not a valid timezone string.", args[0]))
	}
	cache.SetUserLocation(ctx.userID, location)
	ctx.reply(fmt.Sprintf("Okay, your local timezone has been set to **%v**.", location.String()))
}

func (ctx *discordContext) show() {
	location := cache.GetUserLocation(ctx.userID)
	if location != nil {
		ctx.reply(fmt.Sprintf("Currently configured timezone is **%v**", location.String()))
		return
	}

	ctx.reply("Sorry, I don't have any configured timezone for you. Try `set`.")
}

func (ctx *discordContext) getTime() {
	location := cache.GetUserLocation(ctx.userID)
	if location == nil {
		ctx.reply("Sorry, I don't have any configured timezone for you. Try `set`.")
		return
	}

	now := time.Now()
	local, _ := time.LoadLocation("America/Los_Angeles")
	ctx.reply(fmt.Sprintf("`  UTC time is %v`\n", now.Format(format)))
	ctx.reply(fmt.Sprintf("`LOCAL time is %v`", now.In(local).Format(format)))
}

func (ctx *discordContext) reply(message string) {
	mention := ctx.user.Mention()
	msg := fmt.Sprintf("%v %v", mention, message)
	_, err := ctx.session.ChannelMessageSend(ctx.channelID, msg)
	if err != nil {
		ctx.log().Error().Msgf("error sending message to Discord: %v", err.Error())
	}
}

func (ctx *discordContext) log() *zerolog.Logger {
	return ctx.logCtx
}
