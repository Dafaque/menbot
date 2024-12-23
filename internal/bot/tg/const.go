package tg

const (
	CmdStart = "start"
	CmdHelp  = "h"

	CmdList   = "l"
	CmdAdd    = "a"
	CmdRemove = "rm"

	CmdListUsers  = "lu"
	CmdAddUser    = "au"
	CmdRemoveUser = "rmu"

	CmdAddSuperuser = "asu"
)

const (
	parseModeMarkdown = "MarkdownV2"
)

const (
	helpMessage string = `
*h* \- _show this message_
*l* \- _list roles_
*asu* \- _\<token\> submit yourself as superuser_

*How to tag*:
Start typing \/ and find available role\. Press TAB to continue typing with role mention\.
`

	msgAlreadySuperuser = `
*Superuser commands*:
*a* \- _\<roleName\> add role_
*rm* \- _\<roleName\> remove role_
*lu* \- _\<roleName\> list users_
*au* \- _\<roleName\> \<userName\> add user to role_
*rmu* \- _\<roleName\> \<userName\> remove user from role_
`
)
