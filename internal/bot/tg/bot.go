package tg

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Dafaque/tgment/internal/handlers/handler"
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
	bot.SetMyCommands(&telego.SetMyCommandsParams{
		Commands: []telego.BotCommand{
			{Command: "h", Description: "show help"},
		},
	})
	return &Bot{bot: bot, handlers: handler}
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

func (b *Bot) HandleUpdate(update telego.Update) {
	if update.Message == nil {
		return
	}
	if !strings.HasPrefix(update.Message.Text, "/") {
		return
	}
	cmdPart := strings.SplitN(update.Message.Text, "\n", 2)[0]
	if len(cmdPart) < 2 || len(cmdPart) > 512 {
		return
	}
	command := strings.TrimPrefix(cmdPart, "/")
	if strings.Contains(command, "@") {
		cmdAndArgs := strings.SplitN(command, " ", 2)
		components := strings.SplitN(cmdAndArgs[0], "@", 2)
		command = components[0]
		if len(cmdAndArgs) > 1 {
			command += " " + cmdAndArgs[1]
		}
	}
	msgReqUserTgId := update.Message.From.ID

	cmdWithArgs := strings.Split(command, " ")
	if len(cmdWithArgs) < 1 {
		return
	}
	cmd := cmdWithArgs[0]
	args := cmdWithArgs[1:]

	var response string
	switch cmd {
	case CmdHelp:
		response = b.handleHelp(msgReqUserTgId)
	case CmdList:
		response = b.handleListRoles()
	case CmdAdd:
		roleName, err := parseAddRoleArgs(args)
		if err != nil {
			response = err.Error()
			break
		}
		response = b.handleAddRole(msgReqUserTgId, roleName)
	case CmdRemove:
		roleName, err := parseRemoveRoleArgs(args)
		if err != nil {
			response = err.Error()
			break
		}
		response = b.handleRemoveRole(msgReqUserTgId, roleName)
	case CmdListUsers:
		response = b.handleListRoleUsers(msgReqUserTgId)
	case CmdAddUser:
		roleName, userTgId, err := parseRoleUserArgs(args)
		if err != nil {
			response = err.Error()
			break
		}
		response = b.handleAddRoleUser(msgReqUserTgId, roleName, userTgId)
	case CmdRemoveUser:
		roleName, userTgId, err := parseRoleUserArgs(args)
		if err != nil {
			response = err.Error()
			break
		}
		response = b.handleRemoveRoleUser(msgReqUserTgId, roleName, userTgId)
	case CmdTag:
		roleName, err := parseTagRoleArgs(args)
		if err != nil {
			response = err.Error()
			break
		}
		response = b.handleTagRole(roleName)
	case CmdAddSuperuser:
		token, err := parseAddSuperuserArgs(args)
		if err != nil {
			response = err.Error()
			break
		}
		response = b.handleAddSuperuser(token, msgReqUserTgId)
	}

	b.bot.SendMessage(&telego.SendMessageParams{
		ChatID:    update.Message.Chat.ChatID(),
		Text:      response,
		ParseMode: parseModeMarkdown,
		ReplyParameters: &telego.ReplyParameters{
			MessageID: update.Message.MessageID,
		},
	})
}

func (b *Bot) handleHelp(msgReqUserTgId int64) string {
	msg := helpMessage
	if b.handlers.IsUper(msgReqUserTgId) {
		msg += msgAlreadySuperuser
	}
	return msg
}

func (b *Bot) handleListRoles() string {
	return b.handlers.ListRoles()
}

func (b *Bot) handleAddRole(reqUserTgId int64, roleName string) string {
	return b.handlers.AddRole(reqUserTgId, roleName)
}

func (b *Bot) handleRemoveRole(reqUserTgId int64, roleName string) string {
	return b.handlers.RemoveRole(reqUserTgId, roleName)
}

func (b *Bot) handleListRoleUsers(reqUserTgId int64) string {
	return b.handlers.ListRoleUsers(reqUserTgId)
}

func (b *Bot) handleAddRoleUser(reqUserTgId int64, roleName, userTgId string) string {
	return b.handlers.AddRoleUser(reqUserTgId, roleName, userTgId)
}

func (b *Bot) handleRemoveRoleUser(reqUserTgId int64, roleName, userTgId string) string {
	return b.handlers.RemoveRoleUser(reqUserTgId, roleName, userTgId)
}

func (b *Bot) handleAddSuperuser(token string, userTgId int64) string {
	return b.handlers.AddSuperuser(token, userTgId)
}

func (b *Bot) handleTagRole(roleName string) string {
	return b.handlers.GetUsersByRole(roleName)
}

func parseAddRoleArgs(args []string) (string, error) {
	if len(args) < 1 {
		return "", errors.New("*role name* is required")
	}
	roleName := strings.TrimSpace(args[0])
	if roleName == "" {
		return "", errors.New("*role name* is required")
	}
	return roleName, nil
}

func parseRemoveRoleArgs(args []string) (string, error) {
	if len(args) < 1 {
		return "", errors.New("*role name* is required")
	}
	return args[0], nil
}

func parseRoleUserArgs(args []string) (string, string, error) {
	if len(args) < 2 {
		return "", "", errors.New("*role name* and *user tag* are required")
	}
	roleName := strings.TrimSpace(args[0])
	if roleName == "" {
		return "", "", errors.New("*role name* is required")
	}
	userTgId := strings.TrimSpace(args[1])
	if len(userTgId) < 2 {
		return "", "", errors.New("*user tag* is required")
	}
	return roleName, userTgId, nil
}

func parseAddSuperuserArgs(args []string) (string, error) {
	if len(args) < 1 {
		return "", errors.New("*token* is required")
	}
	token := strings.TrimSpace(args[0])
	if token == "" {
		return "", errors.New("*token* is required")
	}
	return token, nil
}

func parseTagRoleArgs(args []string) (string, error) {
	if len(args) < 1 {
		return "", errors.New("*role name* is required")
	}
	return args[0], nil
}
