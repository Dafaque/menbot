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
	case CmdAddSuperuser:
		token, err := parseAddSuperuserArgs(args)
		if err != nil {
			response = err.Error()
			break
		}
		response = b.handleAddSuperuser(token, msgReqUserTgId)
	default:
		response = b.handleTagRole(cmd)
	}

	if len(response) == 0 {
		log.Println("No response for command", cmd)
		return
	}

	_, err := b.bot.SendMessage(&telego.SendMessageParams{
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
	defer b.updateCommands()
	return b.handlers.AddRole(reqUserTgId, roleName)
}

func (b *Bot) handleRemoveRole(reqUserTgId int64, roleName string) string {
	defer b.updateCommands()
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
