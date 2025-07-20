package tg

const (
	CmdStart = "start"
	CmdHelp  = "help"

	CmdAuthorize = "authorize"

	CmdSubscribe   = "subscribe"
	CmdUnsubscribe = "unsubscribe"

	RoleAll = "all"
)

const (
	parseModeMarkdown = "MarkdownV2"
)

const (
	helpMessage string = `
*help* \- _show this message_
*start* \- _allow bot to use you username_
*authorize* \- _authorize chat\. For chat admins only_
*subscribe \<roleName\>* \- _subscribe to the role\. Available roles in bot commands_
*unsubscribe \<roleName\>* \- _unsubscribe from the role\. Available roles in bot commands_

*How to tag*:
Start typing \/ and find available role\. Press TAB to continue typing with role mention\.
`
)
