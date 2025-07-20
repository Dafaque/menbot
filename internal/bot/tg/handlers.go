package tg

import (
	"log"
	"strings"

	"github.com/mymmrac/telego"
)

func (b *Bot) HandleUpdate(update telego.Update) {
	if update.Message == nil {
		return
	}
	if update.Message.From.IsBot {
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
	// msgReqUserTgId := update.Message.From.ID
	// update.
	cmdWithArgs := strings.Split(command, " ")
	if len(cmdWithArgs) < 1 {
		return
	}
	cmd := cmdWithArgs[0]

	var response string
	switch cmd {
	case CmdHelp:
		response = b.handleHelp()
	case CmdStart:
		response = cleanupOutput(b.handleStart(update.Message))
	case CmdAuthorize:
		response = cleanupOutput(b.handleAuthorize(update.Message))
	case CmdSubscribe:
		response = cleanupOutput(b.handleSubscribe(update.Message, cmdWithArgs[1:]))
	case CmdUnsubscribe:
		response = cleanupOutput(b.handleUnsubscribe(update.Message, cmdWithArgs[1:]))
	default:
		response = cleanupOutput(b.handleTagRole(update.Message, cmd))
	}

	if len(response) == 0 {
		log.Println("No response for command", cmd)
		return
	}

	_, err := b.bot.SendMessage(
		b.ctx,
		&telego.SendMessageParams{
			ChatID:    update.Message.Chat.ChatID(),
			Text:      response,
			ParseMode: parseModeMarkdown,
			ReplyParameters: &telego.ReplyParameters{
				MessageID: update.Message.MessageID,
			},
		})
	if err != nil {
		log.Println("Failed to send message", err)
	}
}

func (b *Bot) handleStart(m *telego.Message) string {
	err := b.handlers.Register(
		b.ctx,
		m.Chat.ChatID().ID,
		m.From.ID,
		m.From.Username,
	)
	if err != nil {
		return "Failed to register"
	}
	return "Registered"
}

func (b *Bot) handleHelp() string {
	return helpMessage
}

func (b *Bot) handleTagRole(m *telego.Message, roleName string) string {
	var users []string
	var err error
	if roleName == RoleAll {
		users, err = b.handlers.AllUsers(
			b.ctx,
			m.Chat.ChatID().ID,
		)
	} else {
		users, err = b.handlers.RoleUsers(
			b.ctx,
			m.Chat.ChatID().ID,
			roleName,
		)
	}
	if err != nil {
		return "Failed to get users"
	}
	str := strings.Builder{}
	for _, user := range users {
		if user == m.From.Username {
			continue
		}
		str.WriteString("@")
		str.WriteString(user)
		str.WriteString(", ")
	}
	result, _ := strings.CutSuffix(str.String(), ", ")
	return result
}

func (b *Bot) handleAuthorize(m *telego.Message) string {
	admins, err := b.bot.GetChatAdministrators(b.ctx, &telego.GetChatAdministratorsParams{
		ChatID: telego.ChatID{ID: m.Chat.ChatID().ID},
	})
	if err != nil {
		log.Println("Failed to get chat administrators", err)
		return "Failed to get chat"
	}
	reqByAdmin := false
	for _, member := range admins {
		if member.MemberUser().ID == m.From.ID {
			reqByAdmin = true
			break
		}
	}
	if !reqByAdmin {
		return "Only admins can authorize chats"
	}
	chatID := m.Chat.ChatID()
	err = b.handlers.NewChat(
		b.ctx,
		chatID.ID,
		m.From.ID,
		m.Chat.Title,
		m.From.Username,
	)
	if err != nil {
		return "Failed to authorize chat"
	}
	return "Authorized"
}

func (b *Bot) handleSubscribe(m *telego.Message, args []string) string {
	if len(args) == 0 {
		return "Usage: /subscribe <roleName>"
	}
	roleName := args[0]
	err := b.handlers.Subscribe(
		b.ctx,
		m.Chat.ChatID().ID,
		m.From.ID,
		roleName,
	)
	if err != nil {
		return "Failed to subscribe"
	}
	return "Subscribed"
}

func (b *Bot) handleUnsubscribe(m *telego.Message, args []string) string {
	if len(args) == 0 {
		return "Usage: /unsubscribe <roleName>"
	}
	roleName := args[0]
	err := b.handlers.Unsubscribe(
		b.ctx,
		m.Chat.ChatID().ID,
		m.From.ID,
		roleName,
	)
	if err != nil {
		return "Failed to unsubscribe"
	}
	return "Unsubscribed"
}
func cleanupOutput(in string) string {
	clean := strings.ReplaceAll(in, "_", "\\_")
	clean = strings.ReplaceAll(clean, ">", "\\>")
	clean = strings.ReplaceAll(clean, "<", "\\<")
	return clean
}
