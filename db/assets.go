package db

import (
	"embed"
	_ "embed"
)

//go:embed tables/*.sql
var Tables embed.FS
