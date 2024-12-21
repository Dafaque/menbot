package tg

const (
	CmdStart = "start"
	CmdHelp  = "h"
	CmdTag   = "t"

	CmdList   = "l"
	CmdAdd    = "a"
	CmdRemove = "rm"

	CmdListUsers  = "lu"
	CmdAddUser    = "au"
	CmdRemoveUser = "rmu"

	CmdAddSuperuser = "asu"
)

var nonSuCmd = []string{CmdHelp, CmdAddSuperuser}

const (
	parseModeMarkdown = "MarkdownV2"
)

const (
	helpMessage string = `
*h* \- _show this message_
*t* \- _tag users assigned to role_
*l* \- _list roles_
*asu* \- _\<token\> submit yourself as superuser_
`

	msgAlreadySuperuser = `

Superuser commands:
*a* \- _\<roleName\> add role_
*rm* \- _\<roleName\> remove role_
*lu* \- _\<roleName\> list users_
*au* \- _\<roleName\> \<userName\> add user to role_
*rmu* \- _\<roleName\> \<userName\> remove user from role_
`
)
