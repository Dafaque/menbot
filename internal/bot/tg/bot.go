package tg

import (
	"fmt"
	"log"
	"os"

	"github.com/Dafaque/mentbot/internal/handlers/handler"
	"github.com/mymmrac/telego"
)

type Bot struct {
	bot      *telego.Bot
	handlers handler.Handler
}

func NewBot(token string, handler handler.Handler) *Bot {
	bot, err := telego.NewBot(token)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	b := &Bot{bot: bot, handlers: handler}
	b.updateCommands()
	return b
}

func (b *Bot) Start() error {
	updates, err := b.bot.UpdatesViaLongPolling(&telego.GetUpdatesParams{
		Timeout: 10,
	})
	if err != nil {
		return err
	}
	go func() {
		for update := range updates {
			b.HandleUpdate(update)
		}
	}()

	return nil
}

func (b *Bot) Stop() {
	b.bot.StopLongPolling()
}

func (b *Bot) updateCommands() {
	commands := []telego.BotCommand{
		{Command: CmdHelp, Description: "Show help"},
		{Command: CmdStart, Description: "Start bot"},
	}
	for _, role := range b.handlers.ListRoleItems() {
		commands = append(commands, telego.BotCommand{Command: role, Description: fmt.Sprintf("Tag @%s users", role)})
	}
	err := b.bot.SetMyCommands(&telego.SetMyCommandsParams{
		Commands: commands,
	})
	if err != nil {
		log.Println("Failed to update commands", err)
	}
}
