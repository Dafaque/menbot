package tg

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Dafaque/mentbot/internal/handlers/tg"
	"github.com/mymmrac/telego"
)

type Bot struct {
	bot      *telego.Bot
	handlers tg.Handler
	ctx      context.Context
	cancel   context.CancelFunc
}

func NewBot(token string, handler tg.Handler) *Bot {
	bot, err := telego.NewBot(token)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	b := &Bot{bot: bot, handlers: handler}
	b.ctx, b.cancel = context.WithCancel(context.Background())
	b.UpdateCommands()
	return b
}

func (b *Bot) Start() error {
	updates, err := b.bot.UpdatesViaLongPolling(
		b.ctx,
		&telego.GetUpdatesParams{
			Timeout: 10,
		},
		telego.WithLongPollingUpdateInterval(1*time.Second),
		telego.WithLongPollingRetryTimeout(10*time.Second),
	)
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
	b.cancel()
}

func (b *Bot) UpdateCommands() {

	roles, err := b.handlers.RolesForBotCommands(b.ctx)
	if err != nil {
		log.Println("Failed to get roles for bot commands", err)
	}
	for chatID, roles := range roles {
		commands := []telego.BotCommand{
			{Command: CmdHelp, Description: "Show help"},
			{Command: CmdStart, Description: "Start bot"},
			{Command: CmdAuthorize, Description: "Authorize chat"},
			{Command: CmdSubscribe, Description: "Subscribe to role"},
			{Command: CmdUnsubscribe, Description: "Unsubscribe from role"},
			{Command: RoleAll, Description: "Tag all users"},
		}
		for _, role := range roles {
			commands = append(commands, telego.BotCommand{
				Command:     role,
				Description: fmt.Sprintf("Tag @%s users", role),
			})
		}
		err = b.bot.SetMyCommands(b.ctx, &telego.SetMyCommandsParams{
			Commands: commands,
			Scope: &telego.BotCommandScopeChat{
				Type:   "chat",
				ChatID: telego.ChatID{ID: chatID},
			},
		})
		if err != nil {
			log.Println("Failed to update commands", err)
		}
	}

}
